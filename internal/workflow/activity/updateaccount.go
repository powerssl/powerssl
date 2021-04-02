package activity

import "powerssl.dev/sdk/apiserver/api"

const UpdateAccount = "UpdateAccount"

type UpdateAccountParams struct {
	Name        string
	UpdateMask  []string
	ACMEAccount *api.ACMEAccount
}

func (p *UpdateAccountParams) ToKeyVals() []interface{} {
	return []interface{}{
		"Name", p.Name,
		"UpdateMask", p.UpdateMask,
		"ACMEAccount", p.ACMEAccount,
	}
}

type UpdateAccountResults struct {
	ACMEAccount *api.ACMEAccount
}
