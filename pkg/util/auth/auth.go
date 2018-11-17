package auth

import (
	"context"

	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/spf13/viper"
)

var KeyFunc = func(token *stdjwt.Token) (interface{}, error) {
	return []byte(viper.GetString("signing-key")), nil
}

var Method = stdjwt.SigningMethodHS256

func NewSigner(token string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return next(context.WithValue(ctx, jwt.JWTTokenContextKey, token), request)
		}
	}
}
