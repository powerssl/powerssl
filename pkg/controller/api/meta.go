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

type ChallengeType int32

const (
	ChallengeType_TYPE_UNSPECIFIED ChallengeType = 0
	ChallengeType_DNS_01           ChallengeType = 1
	ChallengeType_HTTP_01          ChallengeType = 2
	ChallengeType_TLS_SNI_01       ChallengeType = 3
)

var ChallengeType_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "DNS_01",
	2: "HTTP_01",
	3: "TLS_SNI_01",
}
var ChallengeType_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"DNS_01":           1,
	"HTTP_01":          2,
	"TLS_SNI_01":       3,
}

func (x ChallengeType) String() string {
	return proto.EnumName(ChallengeType_name, int32(x))
}

type Challenge struct {
	Type    ChallengeType
	Details map[string]string
}

type Error struct {
	Message string
}
