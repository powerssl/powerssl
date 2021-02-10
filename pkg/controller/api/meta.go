package api // import "powerssl.dev/powerssl/pkg/controller/api"

//go:generate gobin -m -run golang.org/x/tools/cmd/stringer -type=ActivityName -trimprefix=Activity
type ActivityName uint

const (
	ActivityUnspecified                 ActivityName = 0
	ActivityACMECreateAccount           ActivityName = 101
	ActivityACMECreateAuthorization     ActivityName = 102
	ActivityACMECreateOrder             ActivityName = 103
	ActivityACMEDeactivateAccount       ActivityName = 104
	ActivityACMEDeactivateAuthorization ActivityName = 105
	ActivityACMEFinalizeOrder           ActivityName = 106
	ActivityACMEGetAuthorization        ActivityName = 107
	ActivityACMEGetCertificate          ActivityName = 108
	ActivityACMEGetChallenge            ActivityName = 109
	ActivityACMEGetOrder                ActivityName = 110
	ActivityACMERekeyAccount            ActivityName = 111
	ActivityACMERevokeCertificate       ActivityName = 112
	ActivityACMEUpdateAccount           ActivityName = 113
	ActivityACMEValidateChallenge       ActivityName = 114
	ActivityDNSCreateRecord             ActivityName = 201
	ActivityDNSDeleteRecord             ActivityName = 202
	ActivityDNSVerifyDomain             ActivityName = 203
)

func (x ActivityName) IntegrationKind() string {
	switch {
	case x > 100 && x < 200:
		return "acme"
	case x > 200 && x < 300:
		return "dns"
	default:
		return ""
	}
}

type Activity struct {
	Name      ActivityName
	Signature string
	Token     string
}

func (a *Activity) String() string {
	return a.Token
}
