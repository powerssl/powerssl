// Code generated by sqlc. DO NOT EDIT.
// source: acme_account.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const CreateACMEAccount = `-- name: CreateACMEAccount :one
INSERT INTO acme_accounts (acme_server_id, display_name, title, description, terms_of_service_agreed, contacts)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, acme_server_id, display_name, title, description, terms_of_service_agreed, contacts, account_url, created_at, updated_at, deleted_at
`

type CreateACMEAccountParams struct {
	AcmeServerID         uuid.UUID `db:"acme_server_id" json:"acmeServerID"`
	DisplayName          string    `db:"display_name" json:"displayName"`
	Title                string    `db:"title" json:"title"`
	Description          string    `db:"description" json:"description"`
	TermsOfServiceAgreed bool      `db:"terms_of_service_agreed" json:"termsOfServiceAgreed"`
	Contacts             string    `db:"contacts" json:"contacts"`
}

func (q *Queries) CreateACMEAccount(ctx context.Context, arg CreateACMEAccountParams) (AcmeAccount, error) {
	row := q.db.QueryRow(ctx, CreateACMEAccount,
		arg.AcmeServerID,
		arg.DisplayName,
		arg.Title,
		arg.Description,
		arg.TermsOfServiceAgreed,
		arg.Contacts,
	)
	var i AcmeAccount
	err := row.Scan(
		&i.ID,
		&i.AcmeServerID,
		&i.DisplayName,
		&i.Title,
		&i.Description,
		&i.TermsOfServiceAgreed,
		&i.Contacts,
		&i.AccountUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const DeleteACMEAccount = `-- name: DeleteACMEAccount :exec
UPDATE acme_accounts
SET deleted_at = NOW()
WHERE deleted_at IS NULL
  AND id = $1
`

func (q *Queries) DeleteACMEAccount(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, DeleteACMEAccount, id)
	return err
}

const GetACMEAccount = `-- name: GetACMEAccount :one
SELECT acme_accounts.id, acme_accounts.acme_server_id, acme_accounts.display_name, acme_accounts.title, acme_accounts.description, acme_accounts.terms_of_service_agreed, acme_accounts.contacts, acme_accounts.account_url, acme_accounts.created_at, acme_accounts.updated_at, acme_accounts.deleted_at
FROM acme_accounts
         INNER JOIN acme_servers
                    ON acme_servers.deleted_at IS NULL
                        AND acme_servers.id = acme_accounts.acme_server_id
WHERE acme_accounts.deleted_at IS NULL
  AND acme_accounts.id = $1
LIMIT 1
`

func (q *Queries) GetACMEAccount(ctx context.Context, id uuid.UUID) (AcmeAccount, error) {
	row := q.db.QueryRow(ctx, GetACMEAccount, id)
	var i AcmeAccount
	err := row.Scan(
		&i.ID,
		&i.AcmeServerID,
		&i.DisplayName,
		&i.Title,
		&i.Description,
		&i.TermsOfServiceAgreed,
		&i.Contacts,
		&i.AccountUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const ListACMEAccounts = `-- name: ListACMEAccounts :many
SELECT id, acme_server_id, display_name, title, description, terms_of_service_agreed, contacts, account_url, created_at, updated_at, deleted_at
FROM acme_accounts
WHERE deleted_at IS NULL
ORDER BY CASE
             WHEN $1::text = 'created_at' THEN created_at
             WHEN $1::text = 'updated_at' THEN updated_at
             END
LIMIT $3 OFFSET $2
`

type ListACMEAccountsParams struct {
	SqlOrder  string `db:"sql_order" json:"sqlOrder"`
	SqlOffset int32  `db:"sql_offset" json:"sqlOffset"`
	SqlLimit  int32  `db:"sql_limit" json:"sqlLimit"`
}

func (q *Queries) ListACMEAccounts(ctx context.Context, arg ListACMEAccountsParams) ([]AcmeAccount, error) {
	rows, err := q.db.Query(ctx, ListACMEAccounts, arg.SqlOrder, arg.SqlOffset, arg.SqlLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AcmeAccount
	for rows.Next() {
		var i AcmeAccount
		if err := rows.Scan(
			&i.ID,
			&i.AcmeServerID,
			&i.DisplayName,
			&i.Title,
			&i.Description,
			&i.TermsOfServiceAgreed,
			&i.Contacts,
			&i.AccountUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListACMEAccountsByParent = `-- name: ListACMEAccountsByParent :many
SELECT acme_accounts.id, acme_accounts.acme_server_id, acme_accounts.display_name, acme_accounts.title, acme_accounts.description, acme_accounts.terms_of_service_agreed, acme_accounts.contacts, acme_accounts.account_url, acme_accounts.created_at, acme_accounts.updated_at, acme_accounts.deleted_at
FROM acme_accounts
         INNER JOIN acme_servers
                    ON acme_servers.deleted_at IS NULL
                        AND acme_servers.id = acme_accounts.acme_server_id
WHERE acme_accounts.deleted_at IS NULL
  AND acme_accounts.acme_server_id = $1
ORDER BY CASE
             WHEN $2::text = 'created_at' THEN acme_accounts.created_at
             WHEN $2::text = 'updated_at' THEN acme_accounts.updated_at
             END
LIMIT $4 OFFSET $3
`

type ListACMEAccountsByParentParams struct {
	AcmeServerID uuid.UUID `db:"acme_server_id" json:"acmeServerID"`
	SqlOrder     string    `db:"sql_order" json:"sqlOrder"`
	SqlOffset    int32     `db:"sql_offset" json:"sqlOffset"`
	SqlLimit     int32     `db:"sql_limit" json:"sqlLimit"`
}

func (q *Queries) ListACMEAccountsByParent(ctx context.Context, arg ListACMEAccountsByParentParams) ([]AcmeAccount, error) {
	rows, err := q.db.Query(ctx, ListACMEAccountsByParent,
		arg.AcmeServerID,
		arg.SqlOrder,
		arg.SqlOffset,
		arg.SqlLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AcmeAccount
	for rows.Next() {
		var i AcmeAccount
		if err := rows.Scan(
			&i.ID,
			&i.AcmeServerID,
			&i.DisplayName,
			&i.Title,
			&i.Description,
			&i.TermsOfServiceAgreed,
			&i.Contacts,
			&i.AccountUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateACMEAccount = `-- name: UpdateACMEAccount :one
UPDATE acme_accounts
SET display_name            = CASE
                                  WHEN $1::bool THEN $2::text
                                  ELSE display_name
    END,
    title                   = CASE
                                  WHEN $3::bool THEN $4::text
                                  ELSE title
        END,
    description             = CASE
                                  WHEN $5::bool THEN $6::text
                                  ELSE description
        END,
    terms_of_service_agreed = CASE
                                  WHEN $7::bool THEN $8::text
                                  ELSE terms_of_service_agreed
        END,
    contacts                = CASE
                                  WHEN $9::bool THEN $10::text
                                  ELSE contacts
        END,
    updated_at              = NOW()
WHERE deleted_at IS NULL
  AND id = $11
RETURNING id, acme_server_id, display_name, title, description, terms_of_service_agreed, contacts, account_url, created_at, updated_at, deleted_at
`

type UpdateACMEAccountParams struct {
	SetDisplayName          bool      `db:"set_display_name" json:"setDisplayName"`
	DisplayName             string    `db:"display_name" json:"displayName"`
	SetTitle                bool      `db:"set_title" json:"setTitle"`
	Title                   string    `db:"title" json:"title"`
	SetDescription          bool      `db:"set_description" json:"setDescription"`
	Description             string    `db:"description" json:"description"`
	SetTermsOfServiceAgreed bool      `db:"set_terms_of_service_agreed" json:"setTermsOfServiceAgreed"`
	TermsOfServiceAgreed    string    `db:"terms_of_service_agreed" json:"termsOfServiceAgreed"`
	SetContacts             bool      `db:"set_contacts" json:"setContacts"`
	Contacts                string    `db:"contacts" json:"contacts"`
	ID                      uuid.UUID `db:"id" json:"id"`
}

func (q *Queries) UpdateACMEAccount(ctx context.Context, arg UpdateACMEAccountParams) (AcmeAccount, error) {
	row := q.db.QueryRow(ctx, UpdateACMEAccount,
		arg.SetDisplayName,
		arg.DisplayName,
		arg.SetTitle,
		arg.Title,
		arg.SetDescription,
		arg.Description,
		arg.SetTermsOfServiceAgreed,
		arg.TermsOfServiceAgreed,
		arg.SetContacts,
		arg.Contacts,
		arg.ID,
	)
	var i AcmeAccount
	err := row.Scan(
		&i.ID,
		&i.AcmeServerID,
		&i.DisplayName,
		&i.Title,
		&i.Description,
		&i.TermsOfServiceAgreed,
		&i.Contacts,
		&i.AccountUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}