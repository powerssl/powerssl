package workflow // import "powerssl.io/powerssl/pkg/controller/workflow"

import (
	"context"

	"powerssl.io/powerssl/pkg/controller/api"
)

type Service interface {
	Create(ctx context.Context, workflow *api.Workflow) (*api.Workflow, error)
}
