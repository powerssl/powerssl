package integration

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("integration not found")

var Integrations integrations

func Delete(uuid uuid.UUID) error {
	return Integrations.Delete(uuid)
}

func Get(uuid uuid.UUID) (*Integration, error) {
	return Integrations.Get(uuid)
}

func GetByKind(kind Kind) (*Integration, error) {
	return Integrations.GetByKind(kind)
}

func Init() {
	Integrations.Init()
}

func Put(integration *Integration) {
	Integrations.Put(integration)
}

func WaitByKind(ctx context.Context, kind Kind) (*Integration, error) {
	return Integrations.WaitByKind(ctx, kind)
}

type integrations struct {
	m         map[uuid.UUID]*Integration
	c         chan struct{}
	listeners struct {
		s []chan struct{}

		sync.Mutex
	}

	sync.Once
	sync.RWMutex
}

func (i *integrations) Delete(uuid uuid.UUID) error {
	i.RLock()
	_, ok := i.m[uuid]
	i.RUnlock()
	if !ok {
		return ErrNotFound
	}
	i.Lock()
	delete(i.m, uuid)
	i.Unlock()
	i.notify()
	return nil
}

func (i *integrations) Get(uuid uuid.UUID) (*Integration, error) {
	i.RLock()
	integration, ok := i.m[uuid]
	i.RUnlock()
	if !ok {
		return nil, ErrNotFound
	}
	return integration, nil
}

func (i *integrations) GetByKind(kind Kind) (*Integration, error) {
	i.RLock()
	defer i.RUnlock()
	for _, integration := range i.m {
		if integration.Kind == kind {
			return integration, nil
		}
	}
	return nil, ErrNotFound
}

func (i *integrations) Init() {
	i.Do(func() {
		i.m = make(map[uuid.UUID]*Integration)
		i.c = make(chan struct{})
		i.listeners.s = []chan struct{}{}

		go func() {
			for {
				<-i.c
				i.listeners.Lock()
				for _, c := range i.listeners.s {
					close(c)
				}
				i.listeners.s = nil
				i.listeners.Unlock()
			}
		}()
	})
}

func (i *integrations) Put(integration *Integration) {
	i.Lock()
	i.m[integration.UUID] = integration
	i.Unlock()
	i.notify()
}

func (i *integrations) WaitByKind(ctx context.Context, kind Kind) (*Integration, error) {
	for {
		integration, err := i.GetByKind(kind)
		if err != nil && err != ErrNotFound {
			return nil, err
		}
		if err == ErrNotFound {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-i.wait():
				continue
			}
		}
		return integration, nil
	}
}

func (i *integrations) notify() {
	i.c <- struct{}{}
}

func (i *integrations) wait() chan struct{} {
	c := make(chan struct{})
	i.listeners.Lock()
	i.listeners.s = append(i.listeners.s, c)
	i.listeners.Unlock()
	return c
}

func init() {
	Init()
}
