create extension if not exists "uuid-ossp";

create table acme_accounts (
                               id uuid not null default uuid_generate_v4(),
                               acme_server_id uuid not null references acme_servers,
                               display_name varchar not null,
                               title varchar not null,
                               description varchar not null,
                               terms_of_service_agreed boolean not null,
                               contacts varchar not null,
                               account_url varchar not null,
                               created_at timestamp not null default now(),
                               updated_at timestamp not null default now(),
                               deleted_at timestamp null,
                               primary key (id)
);

create table acme_servers (
                              id uuid not null default uuid_generate_v4(),
                              display_name varchar not null,
                              directory_url varchar not null,
                              integration_name varchar not null,
                              created_at timestamp not null default now(),
                              updated_at timestamp not null default now(),
                              deleted_at timestamp null,
                              primary key (id)
);
