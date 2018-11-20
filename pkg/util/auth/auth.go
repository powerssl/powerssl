package auth

import (
	"context"
	"errors"
	"io/ioutil"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

var Method = stdjwt.SigningMethodRS256

// var Signer = stdjwt.NewSigner("TODO", key, auth.Method, stdjwt.StandardClaims{})

func NewParser(pubKeyPath string) (endpoint.Middleware, error) {
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return nil, err
	}
	verifyKey, err := stdjwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return jwt.NewParser(func(token *stdjwt.Token) (interface{}, error) {
		_, ok := token.Header["kid"]
		if !ok {
			return nil, errors.New("JWT kid header not specified")
		}
		return verifyKey, nil
	}, Method, jwt.StandardClaimsFactory), nil
}

func NewSigner(token string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return next(context.WithValue(ctx, jwt.JWTTokenContextKey, token), request)
		}
	}
}

func ClaimsFromContext(ctx context.Context) *stdjwt.StandardClaims {
	claims, ok := ctx.Value(jwt.JWTClaimsContextKey).(*stdjwt.StandardClaims)
	if !ok {
		return &stdjwt.StandardClaims{}
	}
	return claims
}

func SubjectFromContext(ctx context.Context) string {
	return ClaimsFromContext(ctx).Subject
}
