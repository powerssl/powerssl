package api

import proto "github.com/gogo/protobuf/proto"

type ActivityName int32

const (
	Activity_NAME_UNSPECIFIED       ActivityName = 0
	Activity_CA_AUTHORIZE_DOMAIN    ActivityName = 101
	Activity_CA_REQUEST_CERTIFICATE ActivityName = 102
	Activity_CA_REVOKE_CERTIFICATE  ActivityName = 103
	Activity_CA_VERIFY_DOMAIN       ActivityName = 104
	Activity_DNS_CREATE_RECORD      ActivityName = 201
	Activity_DNS_DELETE_RECORD      ActivityName = 202
	Activity_DNS_VERIFY_DOMAIN      ActivityName = 203
)

var ActivityName_name = map[int32]string{
	0:   "NAME_UNSPECIFIED",
	101: "CA_AUTHORIZE_DOMAIN",
	102: "CA_REQUEST_CERTIFICATE",
	103: "CA_REVOKE_CERTIFICATE",
	104: "CA_VERIFY_DOMAIN",
	201: "DNS_CREATE_RECORD",
	202: "DNS_DELETE_RECORD",
	203: "DNS_VERIFY_DOMAIN",
}
var ActivityName_value = map[string]int32{
	"NAME_UNSPECIFIED":       0,
	"CA_AUTHORIZE_DOMAIN":    101,
	"CA_REQUEST_CERTIFICATE": 102,
	"CA_REVOKE_CERTIFICATE":  103,
	"CA_VERIFY_DOMAIN":       104,
	"DNS_CREATE_RECORD":      201,
	"DNS_DELETE_RECORD":      202,
	"DNS_VERIFY_DOMAIN":      203,
}

func (x ActivityName) String() string {
	return proto.EnumName(ActivityName_name, int32(x))
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
