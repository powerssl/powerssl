package service

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	otgorm "github.com/smacker/opentracing-gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/internal/app/apiserver/acmeserver/meta"
	"powerssl.io/internal/app/apiserver/acmeserver/model"
	"powerssl.io/pkg/apiserver/api"
	controllerclient "powerssl.io/pkg/controller/client"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

func New(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) meta.Service {
	db.AutoMigrate(&model.ACMEServer{})
	var svc meta.Service
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

func NewBasicService(db *gorm.DB, logger log.Logger, client *controllerclient.GRPCClient) meta.Service {
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

	server := model.NewACMEServerFromAPI(acmeServer)
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

	acmeServer, err := model.FindACMEServerByName(name, db)
	if err != nil {
		return err
	}
	return db.Delete(acmeServer).Error
}

func (bs basicService) Get(ctx context.Context, name string) (*api.ACMEServer, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	acmeServer, err := model.FindACMEServerByName(name, db)
	if err != nil {
		return nil, err
	}
	return acmeServer.ToAPI(), nil
}

func (bs basicService) List(ctx context.Context, pageSize int, pageToken string) ([]*api.ACMEServer, string, error) {
	db := otgorm.SetSpanToGorm(ctx, bs.db)

	if pageSize < 1 {
		pageSize = 10
	} else if pageSize > 20 {
		pageSize = 20
	}
	offset := -1
	if pageToken != "" {
		var err error
		offset, err = strconv.Atoi(pageToken)
		if err != nil {
			return nil, "", status.Error(codes.InvalidArgument, "malformed page token")
		}
	}
	var acmeServers model.ACMEServers
	if err := db.Limit(pageSize + 1).Offset(offset).Find(&acmeServers).Error; err != nil {
		return nil, "", err
	}
	var nextPageToken string
	if len(acmeServers) > pageSize {
		acmeServers = acmeServers[:len(acmeServers)-1]
		if offset == -1 {
			nextPageToken = strconv.Itoa(pageSize)
		} else {
			nextPageToken = strconv.Itoa(offset + pageSize)
		}
	}
	return acmeServers.ToAPI(), nextPageToken, nil
}

func (bs basicService) Update(ctx context.Context, name string, acmeServer *api.ACMEServer) (*api.ACMEServer, error) {
	return nil, ErrUnimplemented
}
