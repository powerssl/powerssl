package workflow

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"

	"powerssl.io/powerssl/pkg/controller/api"
)

type Definition interface {
	Kind() string
	Run(ctx context.Context)
}

type Workflow struct {
	UUID uuid.UUID

	Definition
}

func New(definition Definition) *Workflow {
	w := &Workflow{
		UUID:       uuid.New(),
		Definition: definition,
	}
	Workflows.Put(w)
	return w
}

func (w *Workflow) Name() string {
	return fmt.Sprintf("workflows/%s", w.UUID)
}

func (w *Workflow) ToAPI() *api.Workflow {
	return &api.Workflow{
		Name: w.Name(),
	}
}

func (w *Workflow) Execute(ctx context.Context) chan struct{} {
	c := make(chan struct{})
	go w.execute(ctx, c)
	return c
}

func (w *Workflow) execute(ctx context.Context, c chan struct{}) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Workflow")
	defer span.Finish()
	span.SetTag("kind", w.Kind())
	span.SetTag("name", w.Name())
	w.Run(ctx)
	// select {
	// case <-w.Run(ctx):
	// case <-ctx.Done():
	// 	fmt.Println(ctx.Err()) // TODO
	// }
	close(c)
}
