-- Create the new schema
CREATE SCHEMA IF NOT EXISTS education;

CREATE EXTENSION pgcrypto;


CREATE TYPE test_type AS ENUM (
  'graded', 'practice', 'survey', 'diagnostic', 'cognitive', 'personality'
);

CREATE TYPE question_type AS ENUM (
  'multiple-choice', 'rating-scale', 'open-ended', 'true/false'
);

-- Owners Table
CREATE TABLE IF NOT EXISTS education.owners (
    owner_id uuid PRIMARY KEY default gen_random_uuid(),
    owner_name VARCHAR(255) NOT NULL,
		pass_hash BYTEA NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tests Table
CREATE TABLE IF NOT EXISTS education.tests (
    test_id uuid PRIMARY KEY default gen_random_uuid(),
    owner_id uuid NOT NULL,
    test_name VARCHAR(255) NOT NULL,
    description TEXT,
    test_type test_type NOT NULL DEFAULT 'graded',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES education.owners(owner_id) ON DELETE CASCADE
);


-- Questions Table
CREATE TABLE IF NOT EXISTS education.questions (
    question_id uuid PRIMARY KEY default gen_random_uuid(),
		test_id uuid NOT NULL,
    question_text TEXT NOT NULL,
    question_type question_type NOT NULL DEFAULT 'multiple-choice',
    question_order INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (test_id) REFERENCES education.tests(test_id) ON DELETE CASCADE
);

-- Answers Table
CREATE TABLE IF NOT EXISTS education.answers (
    answer_id uuid PRIMARY KEY default gen_random_uuid(),
    question_id uuid NOT NULL,
    answer_text TEXT NOT NULL,
    answer_order INT NOT NULL,
    score_value DECIMAL(5, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (question_id) REFERENCES education.questions(question_id) ON DELETE CASCADE
);

-- Students Table
CREATE TABLE IF NOT EXISTS education.students (
    student_id uuid PRIMARY KEY default gen_random_uuid(),
    student_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    date_of_birth DATE NOT NULL,
    enrollment_date DATE DEFAULT CURRENT_DATE,
    group_id uuid DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Groups Table
CREATE TABLE IF NOT EXISTS education.groups (
    group_id uuid PRIMARY KEY default gen_random_uuid(),
    group_name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Test Results Table
CREATE TABLE IF NOT EXISTS education.test_results (
    result_id uuid PRIMARY KEY default gen_random_uuid(),
    student_id uuid NOT NULL,
    test_id uuid NOT NULL,
    raw_score DECIMAL(5, 2),
    scaled_score DECIMAL(5, 2),
    status ENUM('complete', 'in progress', 'skipped') NOT NULL DEFAULT 'in progress',
    interpretation TEXT,
    taken_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES education.students(student_id) ON DELETE CASCADE,
    FOREIGN KEY (test_id) REFERENCES education.tests(test_id) ON DELETE CASCADE
);

create table if not exists users (
    id uuid primary key default gen_random_uuid(),
    first_name varchar(32) not null, -- имя
    middle_name varchar(32) not null, -- отчество
    last_name varchar(32) not null, -- фамилия
    email varchar(32) unique not null,
    pass_hash bytea unique not null,
    gender bool, -- role student
    role_id integer references roles(id),
    school_id uuid references school(id),
    classes_id uuid references classes(id),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null
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

create table if not exists discipline (
    id uuid primary key default gen_random_uuid(),
    name varchar(128) not null
);

create table if not exists theme (
    id uuid primary key default gen_random_uuid(),
    name varchar(128) not null,
    discipline_id uuid references discipline(id) not null
);

create table if not exists study_group (
    id uuid primary key default gen_random_uuid(),
    teacher_id uuid references users(id) not null,
    discipline_id uuid references discipline(id) not null,
    class_id uuid references classes(id) not null
);

create table if not exists lesson (
    id uuid primary key default gen_random_uuid(),
    theme_id uuid references theme(id) not null,
    group_id uuid references classes(id) not null,
    teacher_id uuid references users(id) not null
);
