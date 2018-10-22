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

var (
	lock         = sync.Mutex{}
	unknownError = errors.New("Unknown error")
)

type Integration struct {
	Kind apiv1.IntegrationKind
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

type Integrations map[uuid.UUID]Integration

type integrationServiceServer struct {
	integrations Integrations
	logger       log.Logger
}

func New(logger log.Logger, duration metrics.Histogram, integrations Integrations) resource.Resource {
	return &integrationServiceServer{
		integrations: integrations,
		logger:       logger,
	}

}

func (s *integrationServiceServer) RegisterGRPCServer(baseServer *grpc.Server) {
	apiv1.RegisterIntegrationServiceServer(baseServer, s)
}

func (s *integrationServiceServer) Register(request *apiv1.RegisterIntegrationRequest, stream apiv1.IntegrationService_RegisterServer) error {
	integration := Integration{
		activity:   make(chan *apiv1.Activity),
		disconnect: make(chan error),
		Kind:       request.GetKind(),
		Name:       request.GetName(),
		UUID:       uuid.New(),
	}
	s.register(&integration)

	for {
		select {
		case <-stream.Context().Done():
			s.unregister(&integration)
			return stream.Context().Err()
		case activity := <-integration.activity:
			if err := stream.Send(activity); err != nil {
				s.logger.Log("err", err)
				s.unregister(&integration)
				return unknownError
			}
		case err := <-integration.disconnect:
			s.unregister(&integration)
			return err
		}
	}
}

func (s *integrationServiceServer) register(integration *Integration) {
	s.logger.Log("interation", integration.UUID, "action", "connect", "name", integration.Name, "kind", integration.Kind)

	lock.Lock()
	s.integrations[integration.UUID] = *integration
	lock.Unlock()
}

func (s *integrationServiceServer) unregister(integration *Integration) {
	s.logger.Log("integration", integration.UUID, "action", "disconnect")

	lock.Lock()
	delete(s.integrations, integration.UUID)
	lock.Unlock()
}
