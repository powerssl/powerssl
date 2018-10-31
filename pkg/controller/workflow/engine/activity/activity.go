package activity

import (
	"github.com/google/uuid"

	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/integration"
)

type Activity struct {
	GetRequest   interface{}
	SetResponse  interface{}
	UUID         uuid.UUID
	activityName api.ActivityName
}

func New(activityName api.ActivityName) *Activity {
	a := &Activity{
		UUID:         uuid.New(),
		activityName: activityName,
	}
	Activities.Put(a)
	return a
}

func (a *Activity) Execute(integ *integration.Integration) {
	integ.Send(&apiv1.Activity{
		Name:      apiv1.Activity_Name(a.activityName),
		Signature: uuid.New().String(), // TODO
		Token:     a.UUID.String(),
		Workflow: &apiv1.Activity_Workflow{
			Activities: []string{"foo", "bar", "baz"},
		},
	})
}

func (a *Activity) IntegrationKind() integration.IntegrationKind {
	// TODO
	return integration.IntegrationKindACME
}
