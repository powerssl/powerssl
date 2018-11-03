package workflow

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"

	"powerssl.io/pkg/controller/api"
	"powerssl.io/pkg/controller/workflow/engine/activity"
)

type WorkflowInterface interface {
	Run()
}

type Workflow struct {
	Activities []*activity.Activity
	Kind       string
	UUID       uuid.UUID
	ctx        context.Context
	span       opentracing.Span
}

func New(ctx context.Context, kind string) *Workflow {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Workflow")

	w := &Workflow{
		Kind: kind,
		UUID: uuid.New(),
		ctx:  ctx,
		span: span,
	}

	span.SetTag("kind", w.Kind)
	span.SetTag("Name", w.String())

	Workflows.Put(w)
	return w
}

func (w *Workflow) String() string {
	return fmt.Sprintf("workflows/%s", w.UUID)
}

func (w *Workflow) API() *api.Workflow {
	return &api.Workflow{
		Name: w.String(),
		Kind: w.Kind,
	}
}

func (w *Workflow) AddActivity(activity *activity.Activity) {
	w.Activities = append(w.Activities, activity)
}

func (w *Workflow) Execute() {
	go w.execute()
}

func (w *Workflow) execute() {
	for _, activity := range w.Activities {
		<-activity.Execute(w.ctx)
	}
	go func() error {
		select {
		case <-w.ctx.Done():
			return w.ctx.Err()
		}
	}()
	w.finish()
}

func (w *Workflow) finish() {
	w.span.Finish()
}
