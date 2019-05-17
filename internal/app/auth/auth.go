//go:generate go-bindata -pkg auth -prefix ../../../web/auth ../../../web/auth

package auth

import (
	"context"
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	htmltemplate "html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/arschles/go-bindata-html-template"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/log"
	"github.com/pborman/uuid"
	"golang.org/x/sync/errgroup"
	"gopkg.in/square/go-jose.v2"

	"powerssl.io/internal/pkg/auth"
	"powerssl.io/internal/pkg/util"
)

func Run(httpAddr, metricsAddr, jwtPrivateKeyFile, webappURI string) {
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
		return ServeHTTP(ctx, httpAddr, log.With(logger, "component", "http"), jwtPrivateKeyFile, webappURI)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}

func jwksEndpoint(signKeys ...*rsa.PrivateKey) (func(w http.ResponseWriter, req *http.Request), error) {
	jsonWebKeys := make([]jose.JSONWebKey, len(signKeys))
	for i, signKey := range signKeys {
		jsonWebKey := jose.JSONWebKey{
			Algorithm: "RS256",
			Key:       signKey,
			Use:       "sig",
		}
		publicJSONWebKey := jsonWebKey.Public()
		thumbprint, err := publicJSONWebKey.Thumbprint(crypto.SHA1)
		if err != nil {
			return nil, err
		}
		publicJSONWebKey.KeyID = base64.URLEncoding.EncodeToString(thumbprint)
		jsonWebKeys[i] = publicJSONWebKey
	}
	jwks, err := json.Marshal(&jose.JSONWebKeySet{Keys: jsonWebKeys})
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/jwk+json")
		fmt.Fprintln(w, string(jwks))
	}, nil
}

func ServeHTTP(ctx context.Context, addr string, logger log.Logger, jwtPrivateKeyFile, webappURI string) error {
	signBytes, err := ioutil.ReadFile(jwtPrivateKeyFile)
	if err != nil {
		return fmt.Errorf("failed to load signing key %v", err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return fmt.Errorf("failed to load signing key %v", err)
	}

	mux := http.NewServeMux()
	tmpl, err := template.New("index", Asset).Parse("index.html")
	if err != nil {
		return err
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if err := tmpl.Execute(w, map[string]interface{}{"WebAppURI": htmltemplate.URL(webappURI)}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	jwksHandler, err := jwksEndpoint(signKey)
	if err != nil {
		return err
	}
	mux.HandleFunc("/.well-known/jwks.json", jwksHandler)
	mux.HandleFunc("/jwt", func(w http.ResponseWriter, req *http.Request) {
		expiresAt := time.Now().Add(time.Hour * 24).Unix()
		subject := req.URL.Query().Get("sub")
		tokenString, err := generateToken(signKey, subject, expiresAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/jwt")
		fmt.Fprint(w, tokenString)
	})
	mux.HandleFunc("/service", func(w http.ResponseWriter, req *http.Request) {
		key := jose.JSONWebKey{Key: signKey}
		public := key.Public()
		thumbprint, err := public.Thumbprint(crypto.SHA1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token := jwt.NewWithClaims(auth.Method, &jwt.StandardClaims{})
		token.Header["kid"] = base64.URLEncoding.EncodeToString(thumbprint)
		tokenString, err := token.SignedString(signKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/jwt")
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
