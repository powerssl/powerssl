package activity

import "powerssl.io/pkg/controller/api"

func GetRequest(activity *api.Activity) (interface{}, error) {
	a, err := Activities.GetByAPIActivity(activity)
	if err != nil {
		return nil, err
	}
	f, err := a.GetRequest()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GetInput(activity *api.Activity, input interface{}) error {
	a, err := Activities.GetByAPIActivity(activity)
	if err != nil {
		return err
	}
	return a.GetInput(input)
}

func SetResponse(activity *api.Activity) (interface{}, error) {
	a, err := Activities.GetByAPIActivity(activity)
	if err != nil {
		return nil, err
	}
	f, err := a.SetResponse()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func SetResult(activity *api.Activity, result interface{}) error {
	a, err := Activities.GetByAPIActivity(activity)
	if err != nil {
		return err
	}
	return a.SetResult(result)
}
