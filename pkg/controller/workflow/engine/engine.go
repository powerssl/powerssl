package engine

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"

	"powerssl.io/pkg/controller/api"
	"powerssl.io/pkg/controller/integration"
	"powerssl.io/pkg/controller/workflow/engine/activity"
	"powerssl.io/pkg/controller/workflow/engine/workflow"
)

type Engine struct {
	integrations integration.Integrations
	workflows    map[uuid.UUID]workflow.Workflow
	activities   map[uuid.UUID]activity.Activity

	addWorkflowC chan *workflow.Workflow
	addActivityC chan *activity.Activity

	workflowLock sync.RWMutex
	activityLock sync.RWMutex
}

func New(integrations integration.Integrations) *Engine {
	return &Engine{
		integrations: integrations,
		workflows:    make(map[uuid.UUID]workflow.Workflow),
		activities:   make(map[uuid.UUID]activity.Activity),

		addWorkflowC: make(chan *workflow.Workflow),
		addActivityC: make(chan *activity.Activity),

		workflowLock: sync.RWMutex{},
		activityLock: sync.RWMutex{},
	}
}

func (e *Engine) AddWorkflow(workflow *workflow.Workflow) {
	e.addWorkflowC <- workflow
}

func (e *Engine) AddActivity(activity *activity.Activity) {
	e.addActivityC <- activity
}

func (e *Engine) GetActivity(apiactivity *api.Activity) (*activity.Activity, error) {
	uuid, err := apiactivity.UUID()
	if err != nil {
		return nil, err
	}
	e.activityLock.RLock()
	activity, ok := e.activities[uuid]
	e.activityLock.RUnlock()
	if !ok {
		return nil, errors.New("activity not found")
	}
	return &activity, nil
}

func (e *Engine) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case workflow := <-e.addWorkflowC:
			e.addWorkflow(workflow)
		case activity := <-e.addActivityC:
			e.addActivity(activity)
		}
	}
}

func (e *Engine) addWorkflow(workflow *workflow.Workflow) {
	e.workflowLock.Lock()
	e.workflows[workflow.UUID] = *workflow
	e.workflowLock.Unlock()
}

func (e *Engine) removeWorkflow(workflow *workflow.Workflow) {
	e.workflowLock.Lock()
	delete(e.workflows, workflow.UUID)
	e.workflowLock.Unlock()
}

func (e *Engine) addActivity(activity *activity.Activity) {
	e.activityLock.Lock()
	e.activities[activity.UUID] = *activity
	e.activityLock.Unlock()

	integ, err := e.integrations.GetByKind(activity.IntegrationKind())
	if err != nil {
		panic(err) // TODO
	}
	activity.Execute(integ)
}

func (e *Engine) removeActivity(activity *activity.Activity) {
	e.activityLock.Lock()
	delete(e.activities, activity.UUID)
	e.activityLock.Unlock()
}
