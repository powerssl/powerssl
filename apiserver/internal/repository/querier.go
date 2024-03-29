// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateACMEAccount(ctx context.Context, arg CreateACMEAccountParams) (AcmeAccount, error)
	CreateACMEServer(ctx context.Context, arg CreateACMEServerParams) (AcmeServer, error)
	DeleteACMEAccount(ctx context.Context, id uuid.UUID) error
	DeleteACMEServer(ctx context.Context, id uuid.UUID) error
	GetACMEAccount(ctx context.Context, id uuid.UUID) (AcmeAccount, error)
	GetACMEServer(ctx context.Context, id uuid.UUID) (AcmeServer, error)
	ListACMEAccounts(ctx context.Context, arg ListACMEAccountsParams) ([]AcmeAccount, error)
	ListACMEAccountsByParent(ctx context.Context, arg ListACMEAccountsByParentParams) ([]AcmeAccount, error)
	ListACMEServers(ctx context.Context, arg ListACMEServersParams) ([]AcmeServer, error)
	UpdateACMEAccount(ctx context.Context, arg UpdateACMEAccountParams) (AcmeAccount, error)
	UpdateACMEServer(ctx context.Context, arg UpdateACMEServerParams) (AcmeServer, error)
}

var _ Querier = (*Queries)(nil)
