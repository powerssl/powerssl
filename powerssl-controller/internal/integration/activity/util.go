package activity

import (
	"context"
	"errors"

	temporalclient "go.temporal.io/sdk/client"

	"powerssl.dev/sdk/controller/api"
)

func GetInput(ctx context.Context, activity *api.Activity, input interface{}) error {
	a, err := GetByAPIActivity(activity)
	if err != nil {
		return err
	}
	return a.GetInput(ctx, input)
}

func SetResult(ctx context.Context, activity *api.Activity, temporal temporalclient.Client, result interface{}, activityError error) error {
	a, err := GetByAPIActivity(activity)
	if err != nil {
		return err
	}
	return a.SetResult(ctx, temporal, result, activityError)
}

func GetRequestDeprecated(_ *api.Activity) (interface{}, error) {
	return nil, errors.New("DEPRECATED")
}

func SetResponseDeprecated(_ *api.Activity) (interface{}, error) {
	return nil, errors.New("DEPRECATED")
}
