package activity

import (
	"context"
	"encoding/base64"
	"fmt"
	"reflect"

	temporalactivity "go.temporal.io/sdk/activity"
	temporalclient "go.temporal.io/sdk/client"

	apiv1 "powerssl.dev/api/controller/v1"
	"powerssl.dev/sdk/controller/api"

	"powerssl.dev/controller/internal/integration"
)

type Activity struct {
	activityName api.ActivityName
	input        interface{}
	token        string
}

func New(ctx context.Context, activityName api.ActivityName, input interface{}) *Activity {
	activityInfo := temporalactivity.GetInfo(ctx)
	taskToken := activityInfo.TaskToken
	token := base64.RawStdEncoding.EncodeToString(taskToken)

	a := &Activity{
		activityName: activityName,
		input:        input,
		token:        token,
	}

	Put(a)

	return a
}

func (a *Activity) Execute(ctx context.Context) error {
	activityIntegration, err := a.integration(ctx)
	if err != nil {
		return err
	}
	activityIntegration.Send(&apiv1.Activity{
		Name:  apiv1.Activity_Name(a.activityName),
		Token: a.Token(),
		Workflow: &apiv1.Activity_Workflow{
			Activities: []string{},
		},
	})
	return nil
}

func (a *Activity) Token() string {
	return a.token
}

func (a *Activity) GetInput(_ context.Context, v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(a.input))
	return nil
}

func (a *Activity) SetResult(ctx context.Context, temporal temporalclient.Client, result interface{}, activityError error) (err error) {
	var taskToken []byte
	if taskToken, err = base64.RawStdEncoding.DecodeString(a.Token()); err != nil {
		return err
	}
	if err = temporal.CompleteActivity(ctx, taskToken, result, activityError); err != nil {
		return err
	}
	return nil
}

func (a *Activity) integrationKind() integration.Kind {
	return integration.Kind(a.activityName.IntegrationKind())
}

func (a *Activity) integration(ctx context.Context) (*integration.Integration, error) {
	return integration.WaitByKind(ctx, a.integrationKind())
}
