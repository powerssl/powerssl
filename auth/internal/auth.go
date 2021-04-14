package internal

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
	"golang.org/x/sync/errgroup"
	"gopkg.in/square/go-jose.v2"

	"powerssl.dev/backend/auth"
	"powerssl.dev/backend/httpfs"
	"powerssl.dev/common"
	"powerssl.dev/common/log"
	"powerssl.dev/common/transport"

	"powerssl.dev/auth/internal/asset"
	"powerssl.dev/auth/internal/oauth2"
	"powerssl.dev/auth/internal/template"
)

func Run(cfg *Config) (err error) {
	var logger log.Logger
	if logger, err = log.NewLogger(false); err != nil {
		return err
	}
	defer common.ErrWrapSync(logger, &err)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return common.InterruptHandler(ctx, logger)
	})

	if cfg.Metrics.Addr != "" {
		g.Go(func() error {
			return transport.ServeMetrics(ctx, cfg.Metrics.Addr, logger.With("component", "metrics"))
		})
	}

	if cfg.OAuth2.GitHub.ClientID != "" && cfg.OAuth2.GitHub.ClientSecret != "" {
		oauth2.InitGitHubOauth2Config(cfg.Auth.URI, cfg.OAuth2.GitHub.ClientID, cfg.OAuth2.GitHub.ClientSecret)
	}

	g.Go(func() error {
		return ServeHTTP(ctx, cfg.Addr, cfg.Insecure, cfg.TLS.CertFile, cfg.TLS.PrivateKeyFile, logger.With("component", "http"), cfg.JWT.PrivateKeyFile, cfg.WebApp.URI)
	})

	if err = g.Wait(); err != nil {
		switch err.(type) {
		case common.InterruptError:
		default:
			return err
		}
	}
	return nil
}

func ServeHTTP(ctx context.Context, addr string, insecure bool, certFile, keyFile string, logger log.Logger, jwtPrivateKeyFile, webappURI string) (err error) {
	var mux *http.ServeMux
	if mux, err = buildMux(logger, jwtPrivateKeyFile, webappURI); err != nil {
		return err
	}

	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	c := make(chan error)
	go func() {
		if insecure {
			c <- srv.ListenAndServe()
		} else {
			c <- srv.ListenAndServeTLS(certFile, keyFile)
		}
		close(c)
	}()
	logger.Infof("listening on %s", addr)
	select {
	case err = <-c:
		logger.Error(err)
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	case <-ctx.Done():
		logger.Error(ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err = srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
}

func buildMux(logger log.Logger, jwtPrivateKeyFile, webappURI string) (_ *http.ServeMux, err error) {
	var buffer []byte
	{
		tmpl := bindatahtmltemplate.Must(bindatahtmltemplate.New("index", template.Asset).Parse("index.html"))
		data := map[string]interface{}{
			"WebAppURI": webappURI,
		}
		var buf bytes.Buffer
		if err = tmpl.Execute(&buf, data); err != nil {
			return nil, err
		}
		buffer = buf.Bytes()
	}

	signBytes, err := ioutil.ReadFile(jwtPrivateKeyFile)
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
			logger.Error("state is too short")
			http.Error(w, "state is too short", http.StatusInternalServerError)
			return
		}
		provider := stateS[len(stateS)-1]
		var username string
		if username, err = oauth2.UserInfo(context.Background(), provider, state, req.FormValue("code")); err != nil {
			logger.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "username",
			Value:   username,
			Expires: time.Now().Add(24 * time.Hour),
		})
		http.Redirect(w, req, webappURI, http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("/jwt", func(w http.ResponseWriter, req *http.Request) {
		expiresAt := time.Now().Add(time.Hour * 24).Unix()
		subject := req.URL.Query().Get("sub")
		tokenString, err := generateToken(signKey, subject, expiresAt)
		if err != nil {
			logger.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/jwt")
		_, _ = fmt.Fprint(w, tokenString)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		var url string
		if url, err = oauth2.AuthCodeURL(req.FormValue("provider")); err != nil {
			logger.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, "{\"action\":\"redirect\",\"url\":\""+url+"\"}")
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
