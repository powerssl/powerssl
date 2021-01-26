package temporalserver

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/server/common/authorization"
)

var decisionAllow = authorization.Result{Decision: authorization.DecisionAllow}
var decisionDeny = authorization.Result{Decision: authorization.DecisionDeny}

type authorizer struct{}

func newAuthorizer() authorization.Authorizer {
	return &authorizer{}
}

func (a *authorizer) Authorize(_ context.Context, claims *authorization.Claims, target *authorization.CallTarget) (authorization.Result, error) {
	log.Print(fmt.Sprintf("claims: %+v\ntarget: %+v\n", claims, target))

	roles, found := claims.Namespaces[target.Namespace]
	if !found || roles == authorization.RoleUndefined {
		return decisionDeny, nil
	}

	// Allow all operations within "temporal-system" namespace
	// DON'T DO THIS IN A PRODUCTION ENVIRONMENT
	// IN PRODUCTION, only allow calls from properly authenticated and authorized callers
	// We are taking a shortcut in the sample because we don't have TLS or a auth token
	if target.Namespace == "temporal-system" {
		return decisionAllow, nil
	}

	// Allow all calls except UpdateNamespace through when claim mapper isn't invoked
	// Claim mapper is skipped unless TLS is configured or an auth token is passed
	if claims == nil && target.APIName != "UpdateNamespace" {
		return decisionAllow, nil
	}

	// Allow all operations for system-level admins and writers
	if claims.System & (authorization.RoleAdmin | authorization.RoleWriter) != 0 {
		return decisionAllow, nil
	}

	// For other namespaces, deny "UpdateNamespace" API unless the caller has a writer role in it
	if target.APIName == "UpdateNamespace" {
		if claims.Namespaces[target.Namespace] & authorization.RoleWriter != 0 {
			return decisionAllow, nil
		} else {
			return decisionDeny, nil
		}
	}

	return decisionDeny, nil
}





	//if claims == nil {
	//	return decisionDeny, nil
	//}
	//// Check system level permissions
	//if claims.System == authorization.RoleAdmin || claims.System == authorization.RoleWriter {
	//	return decisionAllow, nil
	//}
	//roles, found := claims.Namespaces[target.Namespace]
	//if !found || roles == authorization.RoleUndefined {
	//	return decisionDeny, nil
	//}
	//return decisionAllow, nil

