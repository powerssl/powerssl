package auth

import (
	"context"
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/log"
	"github.com/pborman/uuid"
	"golang.org/x/sync/errgroup"
	"gopkg.in/square/go-jose.v2"

	"powerssl.io/pkg/util"
	"powerssl.io/pkg/util/auth"
)

var templates *template.Template

func init() {
	_, file, _, _ := runtime.Caller(0)
	pattern := filepath.Join(filepath.Dir(file), "templates", "*.tmpl")
	templates = template.Must(template.ParseGlob(pattern))
}

func Run(httpAddr, httpCertFile, httpKeyFile string, httpInsecure bool, metricsAddr, jwtPrivateKeyFile string) {
	logger := util.NewLogger(os.Stdout)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	if metricsAddr != "" {
		g.Go(func() error {
			return util.ServeMetrics(ctx, metricsAddr, log.With(logger, "component", "metrics"))
		})
	}

	g.Go(func() error {
		return ServeHTTP(ctx, httpAddr, log.With(logger, "component", "http"), jwtPrivateKeyFile)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}

func ServeHTTP(ctx context.Context, addr string, logger log.Logger, jwtPrivateKeyFile string) error {
	signBytes, err := ioutil.ReadFile(jwtPrivateKeyFile)
	if err != nil {
		return fmt.Errorf("Failed to load signing key %v", err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return fmt.Errorf("Failed to load signing key %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.tmpl", struct{}{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/.well-known/jwks.json", func(w http.ResponseWriter, req *http.Request) {
		jsonWebKey := jose.JSONWebKey{
			Key: signKey,
		}
		publicJSONWebKey := jsonWebKey.Public()
		thumbprint, err := publicJSONWebKey.Thumbprint(crypto.SHA1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		publicJSONWebKey.Algorithm = "RS256"
		publicJSONWebKey.Use = "sig"
		publicJSONWebKey.KeyID = base64.URLEncoding.EncodeToString(thumbprint)
		jsonWebKeySet := &jose.JSONWebKeySet{
			Keys: []jose.JSONWebKey{publicJSONWebKey},
		}
		jwks, err := json.Marshal(jsonWebKeySet)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, fmt.Sprintln(string(jwks)))
	})
	mux.HandleFunc("/raw", func(w http.ResponseWriter, req *http.Request) {
		expiresAt := time.Now().Add(time.Hour * 24).Unix()
		subject := req.URL.Query().Get("sub")
		tokenString, err := generateToken(signKey, subject, expiresAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, tokenString)
	})
	mux.HandleFunc("/service", func(w http.ResponseWriter, req *http.Request) {
		token := jwt.NewWithClaims(auth.Method, &jwt.StandardClaims{})
		key := jose.JSONWebKey{
			Key: signKey,
		}
		public := key.Public()
		thumbprint, err := public.Thumbprint(crypto.SHA1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token.Header["kid"] = base64.URLEncoding.EncodeToString(thumbprint)
		tokenString, err := token.SignedString(signKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, tokenString)
	})
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	c := make(chan error)
	go func() {
		c <- srv.ListenAndServe()
		close(c)
	}()
	logger.Log("listening", addr)
	select {
	case err := <-c:
		logger.Log("err", err)
		if err != http.ErrServerClosed {
			return err
		}
		return nil
	case <-ctx.Done():
		logger.Log("err", ctx.Err())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	}
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
