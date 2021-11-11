package integration

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"

	apiv1 "powerssl.dev/api/controller/v1"
)

var ServiceDesc = &apiv1.IntegrationService_ServiceDesc

var errUnknown = errors.New("unknown error")

type Service struct {
	apiv1.UnimplementedIntegrationServiceServer
	ctx    context.Context
	logger *zap.SugaredLogger
}

func New(ctx context.Context, logger *zap.SugaredLogger) *Service {
	return &Service{
		ctx: ctx,
		logger: logger,
	}
}

func (s Service) Register(request *apiv1.RegisterIntegrationRequest, stream apiv1.IntegrationService_RegisterServer) error {
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

func (s *Service) register(integration *Integration) {
	s.logger.Infow("connected integration", "integration", integration.UUID, "name", integration.Name, "kind", integration.Kind)

	Put(integration)
}

func (s *Service) unregister(integration *Integration) {
	s.logger.Infow("disconnected integration", "integration", integration.UUID)

	if err := Delete(integration.UUID); err != nil {
		s.logger.Debug(err)
	}
}