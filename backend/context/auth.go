package context

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	jwt2 "github.com/go-kit/kit/auth/jwt"
)

func ClaimsFromContext(ctx context.Context) *jwt.StandardClaims {
	claims, ok := ctx.Value(jwt2.JWTClaimsContextKey).(*jwt.StandardClaims)
	if !ok {
		return &jwt.StandardClaims{}
	}
	return claims
}

func SubjectFromContext(ctx context.Context) string {
	return ClaimsFromContext(ctx).Subject
}

func IsInternal(ctx context.Context) bool {
	return SubjectFromContext(ctx) == "{...}"
}
