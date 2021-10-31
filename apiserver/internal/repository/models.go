// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AcmeAccount struct {
	ID                   uuid.UUID    `db:"id" json:"id"`
	AcmeServerID         uuid.UUID    `db:"acme_server_id" json:"acmeServerID"`
	DisplayName          string       `db:"display_name" json:"displayName"`
	Title                string       `db:"title" json:"title"`
	Description          string       `db:"description" json:"description"`
	TermsOfServiceAgreed bool         `db:"terms_of_service_agreed" json:"termsOfServiceAgreed"`
	Contacts             string       `db:"contacts" json:"contacts"`
	AccountUrl           string       `db:"account_url" json:"accountUrl"`
	CreatedAt            time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt            time.Time    `db:"updated_at" json:"updatedAt"`
	DeletedAt            sql.NullTime `db:"deleted_at" json:"deletedAt"`
}

type AcmeServer struct {
	ID              uuid.UUID    `db:"id" json:"id"`
	DisplayName     string       `db:"display_name" json:"displayName"`
	DirectoryUrl    string       `db:"directory_url" json:"directoryUrl"`
	IntegrationName string       `db:"integration_name" json:"integrationName"`
	CreatedAt       time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time    `db:"updated_at" json:"updatedAt"`
	DeletedAt       sql.NullTime `db:"deleted_at" json:"deletedAt"`
}
