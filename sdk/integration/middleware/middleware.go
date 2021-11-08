package middleware

import (
	apiv1 "powerssl.dev/api/controller/v1"
)

type Handler interface {
	Handle(activity *apiv1.Activity)
}

type Middleware func(Handler) Handler

type activityHandler struct {
	f func(*apiv1.Activity)
}

func ActivityHandler(f func(*apiv1.Activity)) Handler {
	return &activityHandler{
		f: f,
	}
}

func (h *activityHandler) Handle(activity *apiv1.Activity) {
	h.f(activity)
}
