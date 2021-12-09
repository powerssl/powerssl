CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE acme_accounts
(
    id                      UUID      NOT NULL DEFAULT uuid_generate_v4(),
    acme_server_id          UUID      NOT NULL REFERENCES acme_servers,
    display_name            VARCHAR   NOT NULL,
    title                   VARCHAR   NOT NULL,
    description             VARCHAR   NOT NULL,
    terms_of_service_agreed BOOLEAN   NOT NULL,
    contacts                VARCHAR   NOT NULL,
    account_url             VARCHAR   NOT NULL,
    created_at              TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at              TIMESTAMP NULL,
    PRIMARY KEY (id)
);

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
