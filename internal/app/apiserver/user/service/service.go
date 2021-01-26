package service

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-pg/pg/v10"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/internal/app/apiserver/user/model"
	"powerssl.dev/powerssl/pkg/apiserver/api"
	"powerssl.dev/powerssl/pkg/apiserver/user"
)

var ErrUnimplemented = status.Error(codes.Unimplemented, "Coming soon")

func New(db *pg.DB, logger log.Logger) user.Service {
	var svc user.Service
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

func NewBasicService(db *pg.DB, logger log.Logger) user.Service {
	return basicService{
		db:     db,
		logger: logger,
	}
}

func (bs basicService) Create(ctx context.Context, apiuser *api.User) (*api.User, error) {
	user := model.NewUserFromAPI(apiuser)
	if err := bs.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := tx.Model(user).Insert()
		return err
	}); err != nil {
		return nil, err
	}
	return user.ToAPI(), nil
}

func (bs basicService) Delete(ctx context.Context, name string) error {
	user, err := model.FindUserByName(name, bs.db)
	if err != nil {
		return err
	}
	_, err = bs.db.Model(user).Delete()
	return err
}

func (bs basicService) Get(ctx context.Context, name string) (*api.User, error) {
	user, err := model.FindUserByName(name, bs.db)
	if err != nil {
		return nil, err
	}
	return user.ToAPI(), nil
}

func (bs basicService) List(ctx context.Context, pageSize int, pageToken string) ([]*api.User, string, error) {
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
	var users model.Users
	if err := bs.db.Model(&users).Limit(pageSize + 1).Offset(offset).Select(); err != nil {
		return nil, "", err
	}
	var nextPageToken string
	if len(users) > pageSize {
		users = users[:len(users)-1]
		if offset == -1 {
			nextPageToken = strconv.Itoa(pageSize)
		} else {
			nextPageToken = strconv.Itoa(offset + pageSize)
		}
	}
	return users.ToAPI(), nextPageToken, nil
}

func (bs basicService) Update(ctx context.Context, name string, user *api.User) (*api.User, error) {
	return nil, ErrUnimplemented
}
