create extension if not exists "uuid-ossp";

create table acme_servers (
    id uuid default uuid_generate_v4(),
    display_name varchar not null,
    directory_url varchar not null,
    integration_name varchar not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp null,
    primary key (id)
);
