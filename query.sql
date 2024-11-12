-- CreateAnswer
-- name: CreateAnswer :one
INSERT INTO answer (next_answer_id, title)
VALUES ($1, $2)
RETURNING *;

-- GetAnswer
-- name: GetAnswer :one
SELECT *
FROM answer
WHERE id = $1;

-- ListAnswers
-- name: ListAnswers :many
SELECT *
FROM answer;

-- CreateQuestion
-- name: CreateQuestion :one
INSERT INTO question (next_question_id, number, first_answer_id, title)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- GetQuestion
-- name: GetQuestion :one
SELECT *
FROM question
WHERE id = $1;

-- ListQuestions
-- name: ListQuestions :many
SELECT *
FROM question;

-- CreatePsychologicalType
-- name: CreatePsychologicalType :one
INSERT INTO psychological_type (title)
VALUES ($1)
RETURNING *;

-- GetPsychologicalType
-- name: GetPsychologicalType :one
SELECT *
FROM psychological_type
WHERE id = $1;

-- ListPsychologicalTypes
-- name: ListPsychologicalTypes :many
SELECT *
FROM psychological_type;

-- CreatePsychologicalTest
-- name: CreatePsychologicalTest :one
INSERT INTO psychological_test (first_question_id, type_id, owner_id, title)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- GetPsychologicalTest
-- name: GetPsychologicalTest :one
SELECT *
FROM psychological_test
WHERE id = $1;

-- ListPsychologicalTests
-- name: ListPsychologicalTests :many
SELECT *
FROM psychological_test;

-- CreatePsychologicalPerformance
-- name: CreatePsychologicalPerformance :one
INSERT INTO psychological_performance (owner_id, psychological_test_id, started_at)
VALUES ($1, $2, $3)
RETURNING *;

-- GetPsychologicalPerformance
-- name: GetPsychologicalPerformance :one
SELECT *
FROM psychological_performance
WHERE id = $1;

-- ListPsychologicalPerformances
-- name: ListPsychologicalPerformances :many
SELECT *
FROM psychological_performance;

-- CreateUser
-- name: CreateUser :one
INSERT INTO users (email, pass_hash)
VALUES ($1, $2)
RETURNING *;

-- GetUser
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- ListUsers
-- name: ListUsers :many
SELECT *
FROM users;

-- CreateApp
-- name: CreateApp :one
INSERT INTO apps (name, secret)
VALUES ($1, $2)
RETURNING *;

-- GetApp
-- name: GetApp :one
SELECT *
FROM apps
WHERE id = $1;

-- ListApps
-- name: ListApps :many
SELECT *
FROM apps;
