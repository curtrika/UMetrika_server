-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

create table if not exists users (
    id uuid primary key default gen_random_uuid(),
    first_name varchar(32) not null, -- имя
    middle_name varchar(32) not null, -- отчество
    last_name varchar(32) not null, -- фамилия
    email varchar(32) unique not null,
    pass_hash bytea unique not null,
    gender bool, -- role student
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null
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
