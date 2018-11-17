package auth

import (
	"context"
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
	return jwt.NewParser(func(*stdjwt.Token) (interface{}, error) {
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
