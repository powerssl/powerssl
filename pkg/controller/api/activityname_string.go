// Code generated by "stringer -type=ActivityName -trimprefix=Activity"; DO NOT EDIT.

package api

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ActivityUnspecified-0]
	_ = x[ActivityACMECreateAccount-101]
	_ = x[ActivityACMECreateAuthorization-102]
	_ = x[ActivityACMECreateOrder-103]
	_ = x[ActivityACMEDeactivateAccount-104]
	_ = x[ActivityACMEDeactivateAuthorization-105]
	_ = x[ActivityACMEFinalizeOrder-106]
	_ = x[ActivityACMEGetAuthorization-107]
	_ = x[ActivityACMEGetCertificate-108]
	_ = x[ActivityACMEGetChallenge-109]
	_ = x[ActivityACMEGetOrder-110]
	_ = x[ActivityACMERekeyAccount-111]
	_ = x[ActivityACMERevokeCertificate-112]
	_ = x[ActivityACMEUpdateAccount-113]
	_ = x[ActivityACMEValidateChallenge-114]
	_ = x[ActivityDNSCreateRecord-201]
	_ = x[ActivityDNSDeleteRecord-202]
	_ = x[ActivityDNSVerifyDomain-203]
}

const (
	_ActivityName_name_0 = "Unspecified"
	_ActivityName_name_1 = "ACMECreateAccountACMECreateAuthorizationACMECreateOrderACMEDeactivateAccountACMEDeactivateAuthorizationACMEFinalizeOrderACMEGetAuthorizationACMEGetCertificateACMEGetChallengeACMEGetOrderACMERekeyAccountACMERevokeCertificateACMEUpdateAccountACMEValidateChallenge"
	_ActivityName_name_2 = "DNSCreateRecordDNSDeleteRecordDNSVerifyDomain"
)

var (
	_ActivityName_index_1 = [...]uint16{0, 17, 40, 55, 76, 103, 120, 140, 158, 174, 186, 202, 223, 240, 261}
	_ActivityName_index_2 = [...]uint8{0, 15, 30, 45}
)

func (i ActivityName) String() string {
	switch {
	case i == 0:
		return _ActivityName_name_0
	case 101 <= i && i <= 114:
		i -= 101
		return _ActivityName_name_1[_ActivityName_index_1[i]:_ActivityName_index_1[i+1]]
	case 201 <= i && i <= 203:
		i -= 201
		return _ActivityName_name_2[_ActivityName_index_2[i]:_ActivityName_index_2[i+1]]
	default:
		return "ActivityName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}