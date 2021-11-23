package server

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	bindatahtmltemplate "github.com/arschles/go-bindata-html-template"
	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
	"gopkg.in/square/go-jose.v2"

	"powerssl.dev/auth/internal/asset"
	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/template"
	"powerssl.dev/backend/auth"
	"powerssl.dev/backend/httpfs"
	"powerssl.dev/common/log"
)

type Server struct {
	cfg    *Config
	logger log.Logger
	oAuth2 *oauth2.OAuth2
}

func ServeHTTP(ctx context.Context, cfg *Config, logger log.Logger, oAuth2 *oauth2.OAuth2) error {
	return New(cfg, logger, oAuth2).ServeHTTP(ctx)
}

func New(cfg *Config, logger log.Logger, oAuth2 *oauth2.OAuth2) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		oAuth2: oAuth2,
	}
}

func (s *Server) ServeHTTP(ctx context.Context) (err error) {
	var mux *http.ServeMux
	if mux, err = s.buildMux(); err != nil {
		return err
	}

	srv := http.Server{
		Addr:    s.cfg.Addr,
		Handler: mux,
	}

	c := make(chan error)
	go func() {
		if s.cfg.Insecure {
			c <- srv.ListenAndServe()
		} else {
			c <- srv.ListenAndServeTLS(s.cfg.CertFile, s.cfg.KeyFile)
		}
		close(c)
	}()
	s.logger.Infof("listening on %s", s.cfg.Addr)
	select {
	case err = <-c:
		s.logger.Error(err)
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	case <-ctx.Done():
		s.logger.Error(ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err = srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}

func (s *Server) buildMux() (_ *http.ServeMux, err error) {
	var buffer []byte
	{
		tmpl := bindatahtmltemplate.Must(bindatahtmltemplate.New("index", template.Asset).Parse("index.html"))
		data := map[string]interface{}{
			"WebAppURI": s.cfg.WebappURI,
		}
		var buf bytes.Buffer
		if err = tmpl.Execute(&buf, data); err != nil {
			return nil, err
		}
		buffer = buf.Bytes()
	}

	signBytes, err := ioutil.ReadFile(s.cfg.JWTPrivateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load signing key %v", err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to load signing key %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			_, _ = w.Write(buffer)
			return
		}
		if req.URL.Path == "/favicon.ico" {
			w.Header().Add("content-type", "image/x-icon")
			_, _ = w.Write(asset.MustAsset("favicon.ico"))
			return
		}
		http.NotFound(w, req)
	})

	var jwksHandler http.HandlerFunc
	if jwksHandler, err = jwksEndpoint(signKey); err != nil {
		return nil, err
	}
	mux.HandleFunc("/.well-known/jwks.json", jwksHandler)

	mux.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		state := req.FormValue("state")
		stateS := strings.Split(state, ":")
		if len(stateS) < 2 {
			s.logger.Error("state is too short")
			http.Error(w, "state is too short", http.StatusInternalServerError)
			return
		}
		provider := stateS[len(stateS)-1]
		var username *string
		if username, err = s.oAuth2.UserInfo(context.Background(), provider, state, req.FormValue("code")); err != nil {
			s.logger.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "username",
			Value:   *username,
			Expires: time.Now().Add(24 * time.Hour),
		})
		http.Redirect(w, req, s.cfg.WebappURI, http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("/jwt", func(w http.ResponseWriter, req *http.Request) {
		expiresAt := time.Now().Add(time.Hour * 24).Unix()
		subject := req.URL.Query().Get("sub")
		tokenString, err := generateToken(signKey, subject, expiresAt)
		if err != nil {
			s.logger.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/jwt")
		_, _ = fmt.Fprint(w, tokenString)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		var url *string
		if url, err = s.oAuth2.AuthCodeURL(req.FormValue("provider")); err != nil {
			s.logger.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, fmt.Sprintf("{\"action\":\"redirect\",\"url\":\"%s\"}", *url))
	})

	mux.HandleFunc("/logout", func(w http.ResponseWriter, req *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:    "username",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
	})

	mux.HandleFunc("/service", func(w http.ResponseWriter, req *http.Request) {
		key := jose.JSONWebKey{Key: signKey}
		public := key.Public()
		var thumbprint []byte
		if thumbprint, err = public.Thumbprint(crypto.SHA1); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token := jwt.NewWithClaims(auth.Method, &jwt.StandardClaims{})
		token.Header["kid"] = base64.URLEncoding.EncodeToString(thumbprint)
		var tokenString string
		if tokenString, err = token.SignedString(signKey); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/jwt")
		_, _ = fmt.Fprint(w, tokenString)
	})

	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(httpfs.NewFileSystem(asset.AssetFile()))))

	return mux, nil
}

func generateToken(signKey *rsa.PrivateKey, subject string, expiresAt int64) (string, error) {
	claims := &jwt.StandardClaims{
		Audience:  "powerssl.apiserver",
		ExpiresAt: expiresAt,
		Id:        base64.URLEncoding.EncodeToString(uuid.NewRandom())[:22],
		IssuedAt:  time.Now().Unix(),
		Issuer:    "powerssl.auth",
		NotBefore: time.Now().Unix() - 5,
		Subject:   subject,
	}
	token := jwt.NewWithClaims(auth.Method, claims)
	key := jose.JSONWebKey{
		Key: signKey,
	}
	public := key.Public()
	thumbprint, err := public.Thumbprint(crypto.SHA1)
	if err != nil {
		return "", err
	}
	token.Header["kid"] = base64.URLEncoding.EncodeToString(thumbprint)
	return token.SignedString(signKey)
}
