-- name: DeleteACMEAccount :exec
UPDATE acme_accounts
SET deleted_at = NOW()
WHERE deleted_at IS NULL
  AND id = $1;

-- name: CreateACMEAccount :one
INSERT INTO acme_accounts (acme_server_id, display_name, title, description, terms_of_service_agreed, contacts)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetACMEAccount :one
SELECT acme_accounts.*
FROM acme_accounts
         INNER JOIN acme_servers
                    ON acme_servers.deleted_at IS NULL
                        AND acme_servers.id = acme_accounts.acme_server_id
WHERE acme_accounts.deleted_at IS NULL
  AND acme_accounts.id = $1
LIMIT 1;

-- name: ListACMEAccounts :many
SELECT *
FROM acme_accounts
WHERE deleted_at IS NULL
ORDER BY CASE
             WHEN @sql_order::text = 'created_at' THEN created_at
             WHEN @sql_order::text = 'updated_at' THEN updated_at
             END
LIMIT @sql_limit OFFSET @sql_offset;

-- name: ListACMEAccountsByParent :many
SELECT acme_accounts.*
FROM acme_accounts
         INNER JOIN acme_servers
                    ON acme_servers.deleted_at IS NULL
                        AND acme_servers.id = acme_accounts.acme_server_id
WHERE acme_accounts.deleted_at IS NULL
  AND acme_accounts.acme_server_id = @acme_server_id
ORDER BY CASE
             WHEN @sql_order::text = 'created_at' THEN acme_accounts.created_at
             WHEN @sql_order::text = 'updated_at' THEN acme_accounts.updated_at
             END
LIMIT @sql_limit OFFSET @sql_offset;

-- name: UpdateACMEAccount :one
UPDATE acme_accounts
SET display_name            = CASE
                                  WHEN @set_display_name::bool THEN @display_name::text
                                  ELSE display_name
    END,
    title                   = CASE
                                  WHEN @set_title::bool THEN @title::text
                                  ELSE title
        END,
    description             = CASE
                                  WHEN @set_description::bool THEN @description::text
                                  ELSE description
        END,
    terms_of_service_agreed = CASE
                                  WHEN @set_terms_of_service_agreed::bool THEN @terms_of_service_agreed::text
                                  ELSE terms_of_service_agreed
        END,
    contacts                = CASE
                                  WHEN @set_contacts::bool THEN @contacts::text
                                  ELSE contacts
        END,
    updated_at              = NOW()
WHERE deleted_at IS NULL
  AND id = @id
RETURNING *;
