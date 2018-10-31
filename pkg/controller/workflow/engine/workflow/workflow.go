package workflow

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"powerssl.io/pkg/controller/api"
	"powerssl.io/pkg/controller/workflow/engine/activity"
)

var ErrNotFound = errors.New("workflow not found")

type workflows struct {
	m map[uuid.UUID]*Workflow
	sync.Once
	sync.RWMutex
}

func (w *workflows) Delete(uuid uuid.UUID) error {
	w.RLock()
	_, ok := w.m[uuid]
	w.RUnlock()
	if !ok {
		return ErrNotFound
	}
	w.Lock()
	delete(w.m, uuid)
	w.Unlock()
	return nil
}

func (w *workflows) Get(uuid uuid.UUID) (*Workflow, error) {
	w.RLock()
	workflow, ok := w.m[uuid]
	w.RUnlock()
	if !ok {
		return nil, ErrNotFound
	}
	return workflow, nil
}

func (w *workflows) Init() {
	a.Do(func() {
		w.m = make(map[uuid.UUID]*Workflow)
	})
}

func (w *workflows) Put(workflow *Workflow) {
	w.Lock()
	w.m[workflow.UUID] = workflow
	w.Unlock()
}

var Workflows workflows

func init() {
	Workflows.Init()
}

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
