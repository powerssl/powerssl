package api // import "powerssl.dev/powerssl/pkg/controller/api"

import (
	"github.com/google/uuid"
)

type ActivityName uint

const (
	Activity_NAME_UNSPECIFIED              ActivityName = 0
	Activity_ACME_CREATE_ACCOUNT           ActivityName = 101
	Activity_ACME_CREATE_AUTHORIZATION     ActivityName = 102
	Activity_ACME_CREATE_ORDER             ActivityName = 103
	Activity_ACME_DEACTIVATE_ACCOUNT       ActivityName = 104
	Activity_ACME_DEACTIVATE_AUTHORIZATION ActivityName = 105
	Activity_ACME_FINALIZE_ORDER           ActivityName = 106
	Activity_ACME_GET_AUTHORIZATION        ActivityName = 107
	Activity_ACME_GET_CERTIFICATE          ActivityName = 108
	Activity_ACME_GET_CHALLENGE            ActivityName = 109
	Activity_ACME_GET_ORDER                ActivityName = 110
	Activity_ACME_REKEY_ACCOUNT            ActivityName = 111
	Activity_ACME_REVOKE_CERTIFICATE       ActivityName = 112
	Activity_ACME_UPDATE_ACCOUNT           ActivityName = 113
	Activity_ACME_VALIDATE_CHALLENGE       ActivityName = 114
	Activity_DNS_CREATE_RECORD             ActivityName = 201
	Activity_DNS_DELETE_RECORD             ActivityName = 202
	Activity_DNS_VERIFY_DOMAIN             ActivityName = 203
)

func (x ActivityName) String() string {
	return activityName_name[uint(x)]
}

var activityName_name = map[uint]string{
	0:   "Unspecified",
	101: "CreateAccount",
	102: "CreateAuthorization",
	103: "CreateOrder",
	104: "DeactivateAccount",
	105: "DeactivateAuthorization",
	106: "FinalizeOrder",
	107: "GetAuthorization",
	108: "GetCertificate",
	109: "GetChallenge",
	110: "GetOrder",
	111: "RekeyAccount",
	112: "RevokeCertificate",
	113: "UpdateAccount",
	114: "ValidateChallenge",
	201: "CreateRecord",
	202: "DeleteRecord",
	203: "VerifyDomain",
}

type Activity struct {
	Token     string
	Name      ActivityName
	Workflow  *Workflow
	Signature string
}

func (a *Activity) String() string {
	return a.Token
}

func (a *Activity) UUID() (uuid.UUID, error) {
	return uuid.Parse(a.Token)
}
