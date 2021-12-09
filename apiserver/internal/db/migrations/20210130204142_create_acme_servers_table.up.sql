CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE acme_servers
(
    id               UUID      NOT NULL DEFAULT uuid_generate_v4(),
    display_name     VARCHAR   NOT NULL,
    directory_url    VARCHAR   NOT NULL,
    integration_name VARCHAR   NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at       TIMESTAMP NULL,
    PRIMARY KEY (id)
);
