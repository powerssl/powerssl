package integration

import (
	"errors"
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/controller/api/v1"
	resource "powerssl.io/pkg/resource"
)

var ErrNotFound = errors.New("integration not found")

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
	a.Do(func() {
		i.m = make(map[uuid.UUID]*Integration)
	})
}

func (i *integrations) Put(integration *Integration) {
	i.Lock()
	i.m[integration.UUID] = integration
	i.Unlock()
}

var Integrations integrations

func init() {
	Integrations.Init()
}

var (
	unknownError = errors.New("Unknown error")
)

type IntegrationKind string

const (
	IntegrationKindACME IntegrationKind = "acme"
	IntegrationKindDNS  IntegrationKind = "dns"
)

type Integration struct {
	Kind IntegrationKind
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

type integrationServiceServer struct {
	logger log.Logger
}

func New(logger log.Logger, duration metrics.Histogram) resource.Resource {
	return &integrationServiceServer{
		logger: logger,
	}

}

func (s *integrationServiceServer) RegisterGRPCServer(baseServer *grpc.Server) {
	apiv1.RegisterIntegrationServiceServer(baseServer, s)
}

func (s *integrationServiceServer) Register(request *apiv1.RegisterIntegrationRequest, stream apiv1.IntegrationService_RegisterServer) error {
	var kind IntegrationKind
	switch request.GetKind() {
	case apiv1.IntegrationKind_ACME:
		kind = IntegrationKindACME
	case apiv1.IntegrationKind_DNS:
		kind = IntegrationKindDNS
	default:
		return errors.New("integration kind not found")
	}

	integration := &Integration{
		activity:   make(chan *apiv1.Activity),
		disconnect: make(chan error),
		Kind:       kind,
		Name:       request.GetName(),
		UUID:       uuid.New(),
	}
	s.register(integration)

	for {
		select {
		case <-stream.Context().Done():
			s.unregister(integration)
			return stream.Context().Err()
		case activity := <-integration.activity:
			if err := stream.Send(activity); err != nil {
				s.logger.Log("err", err)
				s.unregister(integration)
				return unknownError
			}
		case err := <-integration.disconnect:
			s.unregister(integration)
			return err
		}
	}
}

func (s *integrationServiceServer) register(integration *Integration) {
	s.logger.Log("interation", integration.UUID, "action", "connect", "name", integration.Name, "kind", integration.Kind)

	Integrations.Put(integration)
}

func (s *integrationServiceServer) unregister(integration *Integration) {
	s.logger.Log("integration", integration.UUID, "action", "disconnect")

	Integrations.Delete(integration.UUID)
}
