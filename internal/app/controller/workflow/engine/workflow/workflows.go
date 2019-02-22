package workflow

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("workflow not found")

var Workflows workflows

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
	w.Do(func() {
		w.m = make(map[uuid.UUID]*Workflow)
	})
}

func (w *workflows) Put(workflow *Workflow) {
	w.Lock()
	w.m[workflow.UUID] = workflow
	w.Unlock()
}

func init() {
	Workflows.Init()
}
