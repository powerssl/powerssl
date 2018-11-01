package activity

import (
	"context"

	"github.com/google/uuid"
	"github.com/looplab/fsm"

	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/integration"
)

type Status uint

const (
	Pending Status = iota
	Assinged
	Acknowledged
	Finished
)

type Activity struct {
	GetRequestFunc  interface{}
	SetResponseFunc interface{}
	UUID            uuid.UUID
	activityName    api.ActivityName
	Status
	FSM *fsm.FSM
}

const (
	acknowledge  = "acknowledge"
	acknowledged = "acknowledged"
	assign       = "assign"
	assigned     = "assigned"
	finish       = "finish"
	finished     = "finished"
	pending      = "pending"
)

func New(activityName api.ActivityName) *Activity {
	a := &Activity{
		UUID:         uuid.New(),
		activityName: activityName,
		Status:       Pending,
		FSM: fsm.NewFSM(
			pending,
			fsm.Events{
				{Name: acknowledge, Src: []string{assigned}, Dst: acknowledged},
				{Name: assign, Src: []string{pending}, Dst: assigned},
				{Name: finish, Src: []string{acknowledged}, Dst: finished},
			},
			fsm.Callbacks{
				"enter_state":   func(e *fsm.Event) { println(e) },
				"leave_pending": func(e *fsm.Event) { e.Async() },
			},
		),
	}
	Activities.Put(a)
	return a
}

func (a *Activity) Execute() {
	go a.execute()
}

func (a *Activity) execute() {
	integ, err := integration.Integrations.WaitByKind(context.Background(), a.IntegrationKind())
	if err != nil {
		panic(err) // TODO
	}
	err = a.FSM.Event(assign)
	if e, ok := err.(fsm.AsyncError); !ok && e.Err != nil {
		panic(err)
	}
	integ.Send(&apiv1.Activity{
		Name:      apiv1.Activity_Name(a.activityName),
		Signature: uuid.New().String(), // TODO
		Token:     a.UUID.String(),
		Workflow: &apiv1.Activity_Workflow{
			Activities: []string{"foo", "bar", "baz"},
		},
	})
	if err := a.FSM.Transition(); err != nil {
		panic(err)
	}
}

func (a *Activity) IntegrationKind() integration.IntegrationKind {
	// TODO
	return integration.IntegrationKindACME
}

func (a *Activity) GetRequest() (interface{}, error) {
	if err := a.FSM.Event(acknowledge); err != nil {
		return nil, err
	}
	return a.GetRequestFunc, nil
}

func (a *Activity) SetResponse() (interface{}, error) {
	if err := a.FSM.Event(finish); err != nil {
		return nil, err
	}
	return a.SetResponseFunc, nil
}
