package activity

const CreateVaultTransitKey = "CreateVaultTransitKey"

type CreateVaultTransitKeyParams struct {
	Name string
}

func (p *CreateVaultTransitKeyParams) ToKeyVals() []interface{} {
	return []interface{}{
		"Name", p.Name,
	}
}
