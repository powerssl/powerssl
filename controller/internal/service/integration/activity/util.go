package activity

import (
	"context"

	temporalclient "go.temporal.io/sdk/client"

	apiv1 "powerssl.dev/api/controller/v1"
)

func GetInput(ctx context.Context, activity *apiv1.Activity, input interface{}) error {
	a, err := GetByAPIActivity(activity)
	if err != nil {
		return err
	}
	return a.GetInput(ctx, input)
}

func SetResult(ctx context.Context, activity *apiv1.Activity, temporal temporalclient.Client, result interface{}, activityError error) error {
	a, err := GetByAPIActivity(activity)
	if err != nil {
		return err
	}
	return a.SetResult(ctx, temporal, result, activityError)
}
