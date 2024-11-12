CREATE TABLE IF NOT EXISTS answer (
    id SERIAL PRIMARY KEY,
    next_answer_id INT REFERENCES answer(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS question (
    id SERIAL PRIMARY KEY,
    next_question_id INT REFERENCES question(id) ON DELETE SET NULL,
    number INT NOT NULL,
    first_answer_id INT REFERENCES answer(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS psychological_type (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS  psychological_test (
    id SERIAL PRIMARY KEY,
    first_question_id INT REFERENCES question(id) ON DELETE SET NULL,
    type_id INT REFERENCES psychological_type(id) ON DELETE SET NULL,
    owner_id INT NOT NULL,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS psychological_performance (
    id SERIAL PRIMARY KEY,
    owner_id INT NOT NULL,
    psychological_test_id INT REFERENCES psychological_test(id) ON DELETE CASCADE,
    started_at TIMESTAMPTZ NOT NULL
);

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
    deleted_at timestamp default null,
    role_id integer references roles(id) default 1,
    school_id uuid references school(id),
    classes_id uuid references classes(id)
);

create table if not exists apps (
    id serial primary key,
    name varchar(25) not null unique,
    secret varchar(25) not null unique
);

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

create table if not exists roles (
    id serial primary key,
    title varchar(64) not null
);

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