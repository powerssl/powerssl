package activity

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/looplab/fsm"
	"github.com/opentracing/opentracing-go"

	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/integration"
)

const (
	acknowledge  = "acknowledge"
	acknowledged = "acknowledged"
	assign       = "assign"
	assigned     = "assigned"
	finish       = "finish"
	finished     = "finished"
	pending      = "pending"
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
	FSM  *fsm.FSM
	c    chan struct{}
	span opentracing.Span
}

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
				"leave_pending": func(e *fsm.Event) { e.Async() },
			},
		),
	}
	Activities.Put(a)
	return a
}

func (a *Activity) String() string {
	return fmt.Sprintf("activites/%s", a.UUID)
}

func (a *Activity) Execute(ctx context.Context) chan struct{} {
	a.c = make(chan struct{})
	go a.execute(ctx)
	return a.c
}

func (a *Activity) execute(ctx context.Context) {
	a.span, _ = opentracing.StartSpanFromContext(ctx, "Activity")
	a.span.SetTag("Name", a.activityName)
	a.span.SetTag("UUID", a.UUID)

	err := a.FSM.Event(assign)
	if e, ok := err.(fsm.AsyncError); !ok && e.Err != nil {
		panic(err)
	}

	tmc := opentracing.TextMapCarrier{}
	a.span.Tracer().Inject(a.span.Context(), opentracing.TextMap, tmc)
	textMapCarrier, err := json.Marshal(tmc)
	if err != nil {
		panic(err)
	}

	a.integration().Send(&apiv1.Activity{
		Name:      apiv1.Activity_Name(a.activityName),
		Signature: string(textMapCarrier),
		Token:     a.UUID.String(),
		Workflow: &apiv1.Activity_Workflow{
			Activities: []string{},
		},
	})
	if err := a.FSM.Transition(); err != nil {
		panic(err)
	}
}

func (a *Activity) integrationKind() integration.IntegrationKind {
	// TODO
	return integration.IntegrationKindACME
}

func (a *Activity) integration() *integration.Integration {
	integ, err := integration.Integrations.WaitByKind(context.Background(), a.integrationKind())
	if err != nil {
		// a.FSM.SetState("failed")
		panic(err) // TODO
	}
	return integ
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
	a.span.Finish()
	close(a.c)
	return a.SetResponseFunc, nil
}
