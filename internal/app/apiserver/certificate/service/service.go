package service

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-pg/pg/v10"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/internal/app/apiserver/certificate/model"
	"powerssl.dev/powerssl/pkg/apiserver/api"
	"powerssl.dev/powerssl/pkg/apiserver/certificate"
)

func New(db *pg.DB, logger log.Logger) certificate.Service {
	var svc certificate.Service
	{
		svc = NewBasicService(db, logger)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	db     *pg.DB
	logger log.Logger
}

func NewBasicService(db *pg.DB, logger log.Logger) certificate.Service {
	return basicService{
		db:     db,
		logger: logger,
	}
}

func (bs basicService) Create(ctx context.Context, apicertificate *api.Certificate) (*api.Certificate, error) {
	certificate := model.NewCertificateFromAPI(apicertificate)
	if err := bs.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.Model(certificate).Insert()
		return err
	}); err != nil {
		return nil, err
	}
	return certificate.ToAPI(), nil
	//workflow, err := bs.controllerclient.Workflow.Create(ctx, &controllerapi.Workflow{
	//	Kind: controllerapi.WorkflowKindRequestACMECertificate,
	//	IntegrationFilters: []*controllerapi.WorkflowIntegrationFilter{
	//		&controllerapi.WorkflowIntegrationFilter{},
	//	},
	//	Input: &controllerapi.RequestACMECertificateInput{
	//		DirectoryURL: "https://example.com/directory",
	//		AccountURL:   "https://example.com/acct/123",
	//		Dnsnames:     []string{"example.com", "example.net"},
	//		NotBefore:    "",
	//		NotAfter:     "",
	//	},
	//})
	//if err != nil {
	//	st := status.Convert(err)
	//	bs.logger.Log("err", "rpc", "code", st.Code(), "message", st.Message())
	//	return nil, status.Error(st.Code(), st.Message())
	//}
	//bs.logger.Log("workflow", workflow.Name)
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	certificate, err := model.FindCertificateByName(name, bs.db)
	if err != nil {
		return err
	}
	_, err = bs.db.Model(certificate).Delete()
	return err
}

func (bs basicService) Get(ctx context.Context, name string) (*api.Certificate, error) {
	certificate, err := model.FindCertificateByName(name, bs.db)
	if err != nil {
		return nil, err
	}
	return certificate.ToAPI(), nil
}

func (bs basicService) List(ctx context.Context, pageSize int, pageToken string) ([]*api.Certificate, string, error) {
	var (
		certificates  model.Certificates
		nextPageToken string
	)
	offset := -1
	if pageSize < 1 {
		pageSize = 10
	} else if pageSize > 20 {
		pageSize = 20
	}
	if pageToken != "" {
		var err error
		offset, err = strconv.Atoi(pageToken)
		if err != nil {
			return nil, "", status.Error(codes.InvalidArgument, "malformed page token")
		}
	}
	if err := bs.db.Model(&certificates).Limit(pageSize + 1).Offset(offset).Select(); err != nil {
		return nil, "", err
	}
	if len(certificates) > pageSize {
		certificates = certificates[:len(certificates)-1]
		if offset == -1 {
			nextPageToken = strconv.Itoa(pageSize)
		} else {
			nextPageToken = strconv.Itoa(offset + pageSize)
		}
	}
	return certificates.ToAPI(), nextPageToken, nil
}

func (bs basicService) Update(ctx context.Context, name string, certificate *api.Certificate) (*api.Certificate, error) {
	cert, err := model.FindCertificateByName(name, bs.db)
	if err != nil {
		return nil, err
	}
	if _, err := bs.db.Model(cert).WherePK().Update(model.NewCertificateFromAPI(certificate)); err != nil {
		return nil, err
	}
	cert, err = model.FindCertificateByName(name, bs.db) // TODO Check if needed
	if err != nil {
		return nil, err
	}
	return cert.ToAPI(), nil
}
