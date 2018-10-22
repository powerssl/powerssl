package engine

import (
	"context"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	integration "powerssl.io/pkg/controller/integration"
)

type Engine struct {
	integrations integration.Integrations

	createC chan interface{}
}

func New(integrations integration.Integrations) *Engine {
	return &Engine{
		integrations: integrations,

		createC: make(chan interface{}),
	}
}

func (e *Engine) Create(workflow interface{}) {
	e.createC <- workflow
}

func (e *Engine) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case create := <-e.createC:
			e.create(create)
		}
	}
}

func (e *Engine) create(workflow interface{}) {
	for _, integ := range e.integrations {
		integ.Send(&apiv1.Activity{
			Name:      apiv1.Activity_CA_AUTHORIZE_DOMAIN,
			Signature: "xyz",
			Token:     "token",
			Workflow: &apiv1.Activity_Workflow{
				Activities: []string{"foo", "bar", "baz"},
			},
		})
	}
}
