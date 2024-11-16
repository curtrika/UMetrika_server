-- CreateOwner
-- name: CreateOwner :one
INSERT INTO education.owners (owner_name, email, pass_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- GetOwner
-- name: GetOwner :one
SELECT *
FROM education.owners
WHERE owner_id = $1;

-- ListOwner
-- name: ListOwner :many
SELECT *
FROM education.owners;

-- CreateTest
-- name: CreateTest :one
INSERT INTO education.tests (owner_id, test_name, description, test_type)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- GetTestsByOwnerId
-- name: GetTestsByOwnerId :many
SELECT *
FROM education.tests
WHERE owner_id = $1;
