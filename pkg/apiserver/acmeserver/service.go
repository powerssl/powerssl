package acmeserver

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	otgorm "github.com/smacker/opentracing-gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/pkg/apiserver/api"
	controllerclient "powerssl.io/pkg/controller/client"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

type Service interface {
	Create(ctx context.Context, acmeServer *api.ACMEServer) (*api.ACMEServer, error)
	Delete(ctx context.Context, name string) error
	Get(ctx context.Context, name string) (*api.ACMEServer, error)
	List(ctx context.Context, pageSize int, pageToken string) ([]*api.ACMEServer, string, error)
	Update(ctx context.Context, name string, acmeServer *api.ACMEServer) (*api.ACMEServer, error)
}

func New(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) Service {
	db.AutoMigrate(&ACMEServer{})
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

func (bs basicService) Create(ctx context.Context, acmeServer *api.ACMEServer) (*api.ACMEServer, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return nil, tx.Error
	}

	server := NewACMEServerFromAPI(acmeServer)
	if err := tx.Create(server).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return server.ToAPI(), nil
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	acmeServer, err := FindACMEServerByName(name, db)
	if err != nil {
		return err
	}
	return db.Delete(acmeServer).Error
}

func (bs basicService) Get(ctx context.Context, name string) (*api.ACMEServer, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	acmeServer, err := FindACMEServerByName(name, db)
	if err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (bs basicService) List(ctx context.Context, pageSize int, pageToken string) ([]*api.ACMEServer, string, error) {
	return nil, "", ErrUnimplemented
}

func (bs basicService) Update(ctx context.Context, name string, acmeServer *api.ACMEServer) (*api.ACMEServer, error) {
	return nil, ErrUnimplemented
}
