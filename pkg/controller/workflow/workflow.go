package workflow // import "powerssl.io/pkg/controller/workflow"

import (
	"context"

	"powerssl.io/pkg/controller/api"
)

type Service interface {
	Create(ctx context.Context, workflow *api.Workflow) (*api.Workflow, error)
}
