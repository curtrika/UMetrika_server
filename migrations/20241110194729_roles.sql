-- +goose Up
-- +goose StatementBegin
create table if not exists roles (
    id serial primary key,
    title varchar(64) not null
);

alter table if exists users
    add column if not exists role_id integer references roles(id)
    default 1; -- Ученик
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists users drop column if exists role_id;
drop table if exists roles;
-- +goose StatementEnd