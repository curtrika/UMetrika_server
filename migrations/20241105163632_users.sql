-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

create table if not exists users (
    id uuid primary key default gen_random_uuid(),
    full_name varchar(75) not null,
    email varchar(25) unique not null,
    pass_hash bytea not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz default null
);

create table if not exists apps (
    id serial primary key,
    name varchar(25) not null unique,
    secret varchar(25) not null unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
drop table if exists apps;
-- +goose StatementEnd
