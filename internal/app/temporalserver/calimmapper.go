package temporalserver

import (
	"go.temporal.io/server/common/authorization"
	"go.temporal.io/server/common/service/config"
)

func claimMapper(cfg *config.Config) authorization.ClaimMapper {
	//return authorization.NewNoopClaimMapper(cfg)
	//provider := authorization.NewDefaultTokenKeyProvider(cfg)
	//claimMapper := authorization.NewDefaultJWTClaimMapper(provider, cfg)
	//return claimMapper
	return &myClaimMapper{}
}

type myClaimMapper struct{}

func (c myClaimMapper) GetClaims(authInfo *authorization.AuthInfo) (*authorization.Claims, error) {
	claims := authorization.Claims{}

	if authInfo.TLSConnection != nil {
		claims.Subject = authInfo.TLSSubject.CommonName
	}
	if authInfo.AuthToken != "" {
		// Extract claims from the auth token and translate them into Temporal roles for the caller
		// Here we'll simply hardcode some as an example
		claims.System = authorization.RoleWriter // cluster-level admin
		claims.Namespaces = make(map[string]authorization.Role)
		claims.Namespaces["foo"] = authorization.RoleReader // caller has a reader role for the "foo" namespace
	}

	return &claims, nil
}
