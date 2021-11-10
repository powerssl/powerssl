package activity // import "powerssl.dev/workflow/activity"

import (
	apiv1 "powerssl.dev/api/apiserver/v1"
)

const UpdateAccount = "UpdateAccount"

type UpdateAccountParams struct {
	Name        string
	UpdateMask  []string
	ACMEAccount *apiv1.ACMEAccount
}

func (p *UpdateAccountParams) ToKeyVals() []interface{} {
	return []interface{}{
		"Name", p.Name,
		"UpdateMask", p.UpdateMask,
		"ACMEAccount", p.ACMEAccount,
	}
}

type UpdateAccountResults struct {
	ACMEAccount *apiv1.ACMEAccount
}
