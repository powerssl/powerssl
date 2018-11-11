package acmeaccount

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/pkg/apiserver/api"
	controllerapi "powerssl.io/pkg/controller/api"
	controllerclient "powerssl.io/pkg/controller/client"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

type Service interface {
	Create(ctx context.Context, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.ACMEAccount, error)
	List(ctx context.Context, pageSize int, pageToken string) ([]*api.ACMEAccount, string, error)
	Update(ctx context.Context, name string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error)
}

func New(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) Service {
	db.AutoMigrate(&ACMEAccount{})
	var svc Service
	{
		svc = NewBasicService(db, logger, client)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	controllerclient *controllerclient.GRPCClient
	db               *gorm.DB
	logger           log.Logger
}

func NewBasicService(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) Service {
	return basicService{
		controllerclient: client,
		db:               db,
		logger:           logger,
	}
}

func (bs basicService) Create(ctx context.Context, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	if !acmeAccount.TermsOfServiceAgreed {
		return nil, status.Error(codes.InvalidArgument, "terms of service need to be agreed")
	}

	workflow, err := bs.controllerclient.Workflow.Create(ctx, &controllerapi.Workflow{
		Kind: controllerapi.WorkflowKindCreateACMEAccount,
		IntegrationFilters: []*controllerapi.WorkflowIntegrationFilter{
			&controllerapi.WorkflowIntegrationFilter{},
		},
		Input: &controllerapi.CreateACMEAccountInput{
			DirectoryURL:         "https://example.com/directory",
			TermsOfServiceAgreed: true,
			Contacts:             []string{"mailto:bert@example.com"},
		},
	})
	if err != nil {
		st := status.Convert(err)
		bs.logger.Log("err", "rpc", "code", st.Code(), "message", st.Message())
		return nil, status.Error(st.Code(), st.Message())
	}
	bs.logger.Log("workflow", workflow.Name)
	return acmeAccount, nil
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	return ErrUnimplemented
}

func (bs basicService) Get(ctx context.Context, name string) (*api.ACMEAccount, error) {
	return nil, ErrUnimplemented
}

func (bs basicService) List(ctx context.Context, pageSize int, pageToken string) ([]*api.ACMEAccount, string, error) {
	return nil, "", ErrUnimplemented
}

func (bs basicService) Update(ctx context.Context, name string, acmeAccount *api.ACMEAccount) (*api.ACMEAccount, error) {
	return nil, ErrUnimplemented
}
