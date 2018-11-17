package auth

import (
	"context"
	"crypto/rsa"
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
	"golang.org/x/sync/errgroup"

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
		tokenString, err := generateToken(signKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := templates.ExecuteTemplate(w, "index.tmpl", struct {
			Token string
		}{
			Token: tokenString,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/raw", func(w http.ResponseWriter, req *http.Request) {
		tokenString, err := generateToken(signKey)
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

func generateToken(signKey *rsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(auth.Method, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
	})
	return token.SignedString(signKey)
}
