-- +goose Up

-- Create the new schema
CREATE SCHEMA IF NOT EXISTS education;

-- Enable the pgcrypto extension for UUID generation
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Create custom ENUM types
CREATE TYPE test_type AS ENUM (
  'graded', 'practice', 'survey', 'diagnostic', 'cognitive', 'personality'
);

CREATE TYPE question_type AS ENUM (
  'multiple-choice', 'rating-scale', 'open-ended', 'true/false'
);

CREATE TYPE status AS ENUM (
    'complete', 'in progress', 'skipped'
);


-- Owners Table
CREATE TABLE IF NOT EXISTS education.owners (
    owner_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_name VARCHAR(255) NOT NULL,
    pass_hash BYTEA NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tests Table
CREATE TABLE IF NOT EXISTS education.tests (
    test_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id uuid NOT NULL,
    test_name VARCHAR(255) NOT NULL,
    description TEXT,
    test_type test_type NOT NULL DEFAULT 'graded',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES education.owners(owner_id) ON DELETE CASCADE
);

-- Questions Table
CREATE TABLE IF NOT EXISTS education.questions (
    question_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    test_id uuid NOT NULL,
    question_text TEXT NOT NULL,
    question_type question_type NOT NULL DEFAULT 'multiple-choice',
    question_order INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (test_id) REFERENCES education.tests(test_id) ON DELETE CASCADE
);

-- Answers Table
CREATE TABLE IF NOT EXISTS education.answers (
    answer_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    question_id uuid NOT NULL,
    answer_text TEXT NOT NULL,
    answer_order INT NOT NULL,
    score_value DECIMAL(5, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (question_id) REFERENCES education.questions(question_id) ON DELETE CASCADE
);

-- Students Table
CREATE TABLE IF NOT EXISTS education.students (
    student_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    student_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    date_of_birth DATE NOT NULL,
    enrollment_date DATE DEFAULT CURRENT_DATE,
    group_id uuid DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Groups Table
CREATE TABLE IF NOT EXISTS education.groups (
    group_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    group_name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Test Results Table
CREATE TABLE IF NOT EXISTS education.test_results (
    result_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id uuid NOT NULL,
    test_id uuid NOT NULL,
    raw_score DECIMAL(5, 2),
    scaled_score DECIMAL(5, 2),
    status status NOT NULL DEFAULT 'in progress',
    interpretation TEXT,
    taken_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES education.students(student_id) ON DELETE CASCADE,
    FOREIGN KEY (test_id) REFERENCES education.tests(test_id) ON DELETE CASCADE
);


-- +goose Down

-- Drop all tables and types in reverse order
DROP TABLE IF EXISTS education.test_results CASCADE;
DROP TABLE IF EXISTS education.groups CASCADE;
DROP TABLE IF EXISTS education.students CASCADE;
DROP TABLE IF EXISTS education.answers CASCADE;
DROP TABLE IF EXISTS education.questions CASCADE;
DROP TABLE IF EXISTS education.tests CASCADE;
DROP TABLE IF EXISTS education.owners CASCADE;

DROP TYPE IF EXISTS question_type;
DROP TYPE IF EXISTS test_type;

DROP SCHEMA IF EXISTS education CASCADE;
