package workflow

import (
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
}

func New(kind string) *Workflow {
	return &Workflow{
		Kind: kind,
		UUID: uuid.New(),
	}
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
