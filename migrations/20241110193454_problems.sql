-- +goose Up
-- +goose StatementBegin
create table if not exists problem (
    id uuid primary key default gen_random_uuid(),
    title varchar(1024) not null
);

create table if not exists solution (
    id uuid primary key default gen_random_uuid(),
    title varchar(1024) not null,
    problem_id uuid not null references problem(id)
);

create table if not exists cause (
    id uuid primary key default gen_random_uuid(),
    title varchar(1024) not null,
    problem_id uuid not null references problem(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists cause;
drop table if exists solution;
drop table if exists problem;
-- +goose StatementEnd