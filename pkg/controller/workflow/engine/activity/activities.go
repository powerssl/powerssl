package activity

import (
	"errors"
	"sync"

	"github.com/google/uuid"

	"powerssl.io/pkg/controller/api"
)

var Activities activities

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

func init() {
	Activities.Init()
}
