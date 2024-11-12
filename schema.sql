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
    email varchar(25) unique not null,
    pass_hash bytea not null
);

create table if not exists apps (
    id serial primary key,
    name varchar(25) not null unique,
    secret varchar(25) not null unique
);
