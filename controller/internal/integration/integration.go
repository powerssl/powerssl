package integration

import (
	"github.com/google/uuid"

	apiv1 "powerssl.dev/api/controller/v1"
)

type Kind string

const (
	KindACME Kind = "acme"
	KindDNS  Kind = "dns"
)

type Integration struct {
	Kind Kind
	Name string
	UUID uuid.UUID

	activity   chan *apiv1.Activity
	disconnect chan error
}

func (i *Integration) Disconnect(err error) {
	i.disconnect <- err
}

func (i *Integration) Send(activity *apiv1.Activity) {
	i.activity <- activity
}
