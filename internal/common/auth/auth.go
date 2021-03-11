package auth

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"powerssl.dev/common"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"gopkg.in/square/go-jose.v2"
)

var Method = stdjwt.SigningMethodRS256

// var Signer = stdjwt.NewSigner("TODO", key, auth.Method, stdjwt.StandardClaims{})

func ClaimsFromContext(ctx context.Context) *stdjwt.StandardClaims {
	claims, ok := ctx.Value(jwt.JWTClaimsContextKey).(*stdjwt.StandardClaims)
	if !ok {
		return &stdjwt.StandardClaims{}
	}
	return claims
}

func NewParser(jwksURL string, tlsConfig *tls.Config) (endpoint.Middleware, error) {
	client := http.DefaultClient
	if tlsConfig != nil {
		client = &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}
	}
	resp, err := client.Get(jwksURL)
	if err != nil {
		return nil, err
	}
	defer common.ErrWrapCloser(resp.Body, &err)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var keySet jose.JSONWebKeySet
	if err := json.Unmarshal(body, &keySet); err != nil {
		return nil, err
	}
	return jwt.NewParser(func(token *stdjwt.Token) (interface{}, error) {
		kid, ok := token.Header["kid"]
		if !ok {
			return nil, errors.New("JWT kid header not specified")
		}
		keyID, ok := kid.(string)
		if !ok {
			return nil, errors.New("JWT kid header must be a string")
		}
		keys := keySet.Key(keyID)
		if len(keys) == 0 {
			return nil, errors.New("JWT kid does not match any key")
		}
		return keys[0].Key, nil
	}, Method, jwt.StandardClaimsFactory), nil
}

func NewSigner(token string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return next(context.WithValue(ctx, jwt.JWTTokenContextKey, token), request)
		}
	}
}

func SubjectFromContext(ctx context.Context) string {
	return ClaimsFromContext(ctx).Subject
}
