package activity

import (
	"context"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/looplab/fsm"
	"github.com/opentracing/opentracing-go"

	"powerssl.dev/powerssl/internal/app/controller/integration"
	apiv1 "powerssl.dev/powerssl/internal/pkg/controller/api/v1"
	"powerssl.dev/powerssl/internal/pkg/tracing"
	"powerssl.dev/powerssl/pkg/controller/api"
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

type ActivityInterface interface {
	ActivityName() api.ActivityName
	GetInputType() interface{}
}

type Activity struct {
	FSM             *fsm.FSM
	GetRequestFunc  interface{}
	SetResponseFunc interface{}
	UUID            uuid.UUID
	activityName    api.ActivityName
	c               chan struct{}

	input  interface{}
	result interface{}

	Definition ActivityInterface
}

func NewV2(activity ActivityInterface) *Activity {
	a := New(activity.ActivityName())
	a.Definition = activity
	return a
}

func New(activityName api.ActivityName) *Activity {
	a := &Activity{
		UUID:         uuid.New(),
		activityName: activityName,
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
		c: make(chan struct{}),
	}
	Activities.Put(a)
	return a
}

func (a *Activity) Name() string {
	return fmt.Sprintf("activites/%s", a.UUID)
}

func (a *Activity) Execute(ctx context.Context) chan struct{} {
	c := make(chan struct{})
	go a.execute(ctx, c)
	return c
}

func (a *Activity) execute(ctx context.Context, c chan struct{}) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Activity")
	defer span.Finish()
	span.SetTag("kind", a.activityName)
	span.SetTag("name", a.Name())

	err := a.FSM.Event(assign)
	if e, ok := err.(fsm.AsyncError); !ok && e.Err != nil {
		panic(err)
	}

	carrier, err := tracing.JSONCarrierFromSpan(span)
	if err != nil {
		panic(err)
	}

	a.integration().Send(&apiv1.Activity{
		Name:      apiv1.Activity_Name(a.activityName),
		Signature: carrier,
		Token:     a.UUID.String(),
		Workflow: &apiv1.Activity_Workflow{
			Activities: []string{},
		},
	})
	if err := a.FSM.Transition(); err != nil {
		panic(err)
	}

	select {
	case <-a.c:
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // TODO
	}
	close(c)
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
	close(a.c)
	return a.SetResponseFunc, nil
}

func (a *Activity) GetInput(v interface{}) error {
	if err := a.FSM.Event(acknowledge); err != nil {
		return err
	}
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(a.input)) // TODO: Panic handling
	return nil
}

func (a *Activity) SetInput(input interface{}) {
	// if reflect.TypeOf(a.Definition.GetInputType()) != reflect.TypeOf(input) {
	// 	panic("BOOM")
	// }
	a.input = input
}

func (a *Activity) GetResult(v interface{}) error {
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(a.result)) // TODO: Panic handling
	return nil
}

func (a *Activity) SetResult(result interface{}) error {
	if err := a.FSM.Event(finish); err != nil {
		return err
	}
	close(a.c)
	a.result = result
	return nil
}
