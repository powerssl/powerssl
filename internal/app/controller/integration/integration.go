package integration

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	apiv1 "powerssl.dev/powerssl/internal/pkg/controller/api/v1"
)

var errUnknown = errors.New("unknown error")

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

type integrationServiceServer struct {
	ctx    context.Context
	logger log.Logger
}

func New(ctx context.Context, logger log.Logger, duration metrics.Histogram) *integrationServiceServer {
	return &integrationServiceServer{
		ctx:    ctx,
		logger: logger,
	}

}

func (s *integrationServiceServer) RegisterGRPCServer(baseServer *grpc.Server) {
	apiv1.RegisterIntegrationServiceServer(baseServer, s)
}

func (*integrationServiceServer) ServiceName() string {
	return "powerssl.controller.v1.IntegrationService"
}

func (s *integrationServiceServer) Register(request *apiv1.RegisterIntegrationRequest, stream apiv1.IntegrationService_RegisterServer) error {
	var kind Kind
	switch request.GetKind() {
	case apiv1.IntegrationKind_ACME:
		kind = KindACME
	case apiv1.IntegrationKind_DNS:
		kind = KindDNS
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
		case <-s.ctx.Done():
			s.unregister(integration)
			return s.ctx.Err()
		case <-stream.Context().Done():
			s.unregister(integration)
			return stream.Context().Err()
		case activity := <-integration.activity:
			if err := stream.Send(activity); err != nil {
				s.logger.Log("err", err)
				s.unregister(integration)
				return errUnknown
			}
		case err := <-integration.disconnect:
			s.unregister(integration)
			return err
		}
	}
}

func (s *integrationServiceServer) register(integration *Integration) {
	s.logger.Log("interation", integration.UUID, "action", "connect", "name", integration.Name, "kind", integration.Kind)

	Put(integration)
}

func (s *integrationServiceServer) unregister(integration *Integration) {
	s.logger.Log("integration", integration.UUID, "action", "disconnect")

	Delete(integration.UUID)
}
