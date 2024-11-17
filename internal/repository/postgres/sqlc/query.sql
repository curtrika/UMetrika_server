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

-- InsertQuestionToTest
-- name: InsertQuestionToTest :one
INSERT INTO education.questions (test_id, question_text, question_order)
VALUES ($1, $2, $3)
RETURNING *;

-- InsertAnswerToTest
-- name: InsertAnswerToTest :one
INSERT INTO education.answers (question_id, answer_text, answer_order)
VALUES ($1, $2, $3)
RETURNING *;

-- GetFullTestByOwnerId
-- name: GetFullTestByOwnerId :many
SELECT *
FROM
    education.tests t
JOIN
    education.questions q ON q.test_id = t.test_id
JOIN
    education.answers a ON a.question_id = q.question_id
WHERE t.owner_id = $1;

-- GetTestsByOwnerId
-- name: GetTestsByOwnerId :many
SELECT *
FROM education.tests
WHERE owner_id = $1;

-- GetTeacherDisciplinesAndClasses
-- name: GetTeacherDisciplinesAndClasses :many
SELECT JSON_BUILD_OBJECT(
    'discipline_id', sg.discipline_id,
    'discipline_title', (SELECT name FROM discipline dis WHERE dis.id = sg.discipline_id),
    'classes', (
        SELECT JSON_AGG(
            JSON_BUILD_OBJECT(
                'class_id', c.id,
                'title', concat(c.grade, ' ', c.title),
                'students', (
                    SELECT JSON_AGG(
                        JSON_BUILD_OBJECT(
                            'id', s.id,
                            'first_name', s.first_name,
                            'middle_name', s.middle_name,
                            'last_name', s.last_name
                        )
                    ) FROM users s WHERE s.classes_id = c.id AND s.role_id = 1
                )
            )
        ) FROM classes c
    )
) AS result
FROM study_group sg
WHERE teacher_id = $1;
