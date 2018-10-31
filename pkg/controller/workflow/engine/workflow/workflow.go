package workflow

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"powerssl.io/pkg/controller/api"
	"powerssl.io/pkg/controller/workflow/engine/activity"
)

type WorkflowInterface interface {
	Run()
}

type Workflow struct {
	Kind       string
	UUID       uuid.UUID
	Activities []*activity.Activity
	ctx        context.Context
}

func New(kind string) *Workflow {
	w := &Workflow{
		Kind: kind,
		UUID: uuid.New(),
		ctx:  context.Background(),
	}
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
	go func() error {
		select {
		case <-w.ctx.Done():
			return w.ctx.Err()
		}
	}()
}
