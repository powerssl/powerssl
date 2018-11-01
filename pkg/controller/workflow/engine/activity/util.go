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
