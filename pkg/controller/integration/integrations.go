package integration

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("integration not found")

var Integrations integrations

type integrations struct {
	m map[uuid.UUID]*Integration
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

func (i *integrations) GetByKind(kind IntegrationKind) (*Integration, error) {
	i.RLock()
	defer i.RUnlock()
	for _, integration := range i.m {
		if integration.Kind == kind {
			return integration, nil
		}
	}
	return nil, errors.New("no integration of that type found")
}

func (i *integrations) Init() {
	i.Do(func() {
		i.m = make(map[uuid.UUID]*Integration)
	})
}

func (i *integrations) Put(integration *Integration) {
	i.Lock()
	i.m[integration.UUID] = integration
	i.Unlock()
}

func init() {
	Integrations.Init()
}
