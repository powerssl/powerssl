-- name: DeleteACMEServer :exec
UPDATE acme_servers
SET deleted_at = NOW()
WHERE deleted_at IS NULL
  AND id = $1;

-- name: CreateACMEServer :one
INSERT INTO acme_servers (display_name, directory_url, integration_name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetACMEServer :one
SELECT *
FROM acme_servers
WHERE deleted_at IS NULL
  AND id = $1
LIMIT 1;

-- name: ListACMEServers :many
SELECT *
FROM acme_servers
WHERE deleted_at IS NULL
ORDER BY CASE
             WHEN @sql_order::text = 'created_at' THEN created_at
             WHEN @sql_order::text = 'updated_at' THEN updated_at
             END
LIMIT @sql_limit OFFSET @sql_offset;

-- name: UpdateACMEServer :one
UPDATE acme_servers
SET display_name     = CASE
                           WHEN @set_display_name::bool THEN @display_name::text
                           ELSE display_name
    END,
    directory_url    = CASE
                           WHEN @set_directory_url::bool THEN @directory_url::text
                           ELSE directory_url
        END,
    integration_name = CASE
                           WHEN @set_integration_name::bool THEN @integration_name::text
                           ELSE integration_name
        END,
    updated_at       = NOW()
WHERE deleted_at IS NULL
  AND id = @id
RETURNING *;
