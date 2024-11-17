-- +goose Up
-- +goose StatementBegin
create table if not exists discipline (
    id uuid primary key default gen_random_uuid(),
    name varchar(128) not null
);

create table if not exists theme (
    id uuid primary key default gen_random_uuid(),
    name varchar(128) not null,
    discipline_id uuid references discipline(id) not null
);

create table if not exists lesson (
    id uuid primary key default gen_random_uuid(),
    discipline_id uuid references discipline(id) not null,
    theme_id uuid references theme(id) not null,
    class_id uuid references classes(id) not null,
    teacher_id uuid references users(id) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists lesson;
drop table if exists theme;
drop table if exists discipline;
-- +goose StatementEnd
