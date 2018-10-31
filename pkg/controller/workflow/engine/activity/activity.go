package activity

import (
	"errors"
	"sync"

	"github.com/google/uuid"

	"powerssl.io/pkg/controller/api"
	apiv1 "powerssl.io/pkg/controller/api/v1"
	"powerssl.io/pkg/controller/integration"
)

var ErrNotFound = errors.New("activity not found")

type activities struct {
	m map[uuid.UUID]*Activity
	sync.Once
	sync.RWMutex
}

func (a *activities) Delete(uuid uuid.UUID) error {
	a.RLock()
	_, ok := a.m[uuid]
	a.RUnlock()
	if !ok {
		return ErrNotFound
	}
	a.Lock()
	delete(a.m, uuid)
	a.Unlock()
	return nil
}

func (a *activities) Get(uuid uuid.UUID) (*Activity, error) {
	a.RLock()
	activity, ok := a.m[uuid]
	a.RUnlock()
	if !ok {
		return nil, ErrNotFound
	}
	return activity, nil
}

func (a *activities) GetByAPIActivity(apiactivity *api.Activity) (*Activity, error) {
	uuid, err := apiactivity.UUID()
	if err != nil {
		return nil, err
	}
	return a.Get(uuid)
}

func (a *activities) Init() {
	a.Do(func() {
		a.m = make(map[uuid.UUID]*Activity)
	})
}

func (a *activities) Put(activity *Activity) {
	a.Lock()
	a.m[activity.UUID] = activity
	a.Unlock()
}

var Activities activities

func init() {
	Activities.Init()
}

type Activity struct {
	GetRequest   interface{}
	SetResponse  interface{}
	UUID         uuid.UUID
	activityName api.ActivityName
}

func New(activityName api.ActivityName) *Activity {
	a := &Activity{
		UUID:         uuid.New(),
		activityName: activityName,
	}
	Activities.Put(a)
	return a
}

func (a *Activity) Execute(integ *integration.Integration) {
	integ.Send(&apiv1.Activity{
		Name:      apiv1.Activity_Name(a.activityName),
		Signature: uuid.New().String(), // TODO
		Token:     a.UUID.String(),
		Workflow: &apiv1.Activity_Workflow{
			Activities: []string{"foo", "bar", "baz"},
		},
	})
}

func (a *Activity) IntegrationKind() integration.IntegrationKind {
	// TODO
	return integration.IntegrationKindACME
}
