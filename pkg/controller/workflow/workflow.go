package workflow // import "powerssl.dev/powerssl/pkg/controller/workflow"

import (
	"context"

	"powerssl.dev/powerssl/pkg/controller/api"
)

type Service interface {
	Create(ctx context.Context, workflow *api.Workflow) (*api.Workflow, error)
}
