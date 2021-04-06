package integration

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	apiv1 "powerssl.dev/api/controller/v1"
	backendtransport "powerssl.dev/backend/transport"
	"powerssl.dev/common/log"
)

var errUnknown = errors.New("unknown error")

func NewService(ctx context.Context, logger log.Logger) backendtransport.Service {
	service := &integrationServiceServer{
		ctx:    ctx,
		logger: logger,
	}

	return backendtransport.Service{
		ServiceName: "powerssl.controller.v1.IntegrationService",
		RegisterGRPCServer: func(baseServer *grpc.Server) {
			apiv1.RegisterIntegrationServiceServer(baseServer, service)
		},
	}
}

type integrationServiceServer struct {
	ctx    context.Context
	logger log.Logger
	apiv1.UnimplementedIntegrationServiceServer
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
				s.logger.Error(err)
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
	s.logger.Infow("connected integration", "integration", integration.UUID, "name", integration.Name, "kind", integration.Kind)

	Put(integration)
}

func (s *integrationServiceServer) unregister(integration *Integration) {
	s.logger.Infow("disconnected integration", "integration", integration.UUID)

	if err := Delete(integration.UUID); err != nil {
		s.logger.Debug(err)
	}
}
