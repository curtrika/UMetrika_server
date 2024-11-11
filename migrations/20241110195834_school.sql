-- +goose Up
-- +goose StatementBegin
create table if not exists school (
    id uuid primary key default gen_random_uuid(),
    large_name varchar(1024) not null, -- example Муниципальное бюджетное общеобразовательное...
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp default null
);

create table if not exists classes (
    id uuid primary key default gen_random_uuid(),
    grade int not null default 1,
    title varchar(2),
    main_teacher_id uuid not null references users(id),
    release_date timestamp,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp default null
);

alter table if exists users
    add column if not exists school_id uuid
    references school(id);
alter table if exists users
    add column if not exists classes_id uuid
    references classes(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists users drop column if exists classes_id;
alter table if exists users drop column if exists school_id;
drop table if exists classes;
drop table if exists school;
-- +goose StatementEnd
