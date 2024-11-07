-- +goose Up
-- +goose StatementBegin
create table if not exists roles (
    id uuid primary key default gen_random_uuid(),
    title text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz default null
);

alter table if exists users add column role_id uuid references roles(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists users drop column role_id;
drop table if exists roles;
-- +goose StatementEnd
