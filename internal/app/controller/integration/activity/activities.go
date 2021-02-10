package activity

import (
	"errors"
	"sync"

	"powerssl.dev/powerssl/pkg/controller/api"
)

var Activities activities

var ErrNotFound = errors.New("activity not found")

func Delete(token string) error {
	return Activities.Delete(token)
}

func Get(token string) (*Activity, error) {
	return Activities.Get(token)
}

func GetByAPIActivity(activity *api.Activity) (*Activity, error) {
	return Activities.GetByAPIActivity(activity)
}

func Init() {
	Activities.Init()
}

func Put(activity *Activity) {
	Activities.Put(activity)
}

type activities struct {
	m map[string]*Activity
	sync.Once
	sync.RWMutex
}

func (a *activities) Delete(token string) error {
	a.RLock()
	_, ok := a.m[token]
	a.RUnlock()
	if !ok {
		return ErrNotFound
	}
	a.Lock()
	delete(a.m, token)
	a.Unlock()
	return nil
}

func (a *activities) Get(token string) (*Activity, error) {
	a.RLock()
	activity, ok := a.m[token]
	a.RUnlock()
	if !ok {
		return nil, ErrNotFound
	}
	return activity, nil
}

func (a *activities) GetByAPIActivity(activity *api.Activity) (*Activity, error) {
	return a.Get(activity.Token)
}

func (a *activities) Init() {
	a.Do(func() {
		a.m = make(map[string]*Activity)
	})
}

func (a *activities) Put(activity *Activity) {
	a.Lock()
	a.m[activity.Token()] = activity
	a.Unlock()
}

func init() {
	Init()
}
