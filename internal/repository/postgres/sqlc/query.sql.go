// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOwner = `-- name: CreateOwner :one
INSERT INTO education.owners (owner_name, email, pass_hash)
VALUES ($1, $2, $3)
RETURNING owner_id, owner_name, pass_hash, email, created_at
`

type CreateOwnerParams struct {
	OwnerName string
	Email     string
	PassHash  []byte
}

// CreateOwner
func (q *Queries) CreateOwner(ctx context.Context, arg CreateOwnerParams) (EducationOwner, error) {
	row := q.db.QueryRow(ctx, createOwner, arg.OwnerName, arg.Email, arg.PassHash)
	var i EducationOwner
	err := row.Scan(
		&i.OwnerID,
		&i.OwnerName,
		&i.PassHash,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const createTest = `-- name: CreateTest :one
INSERT INTO education.tests (owner_id, test_name, description, test_type)
VALUES ($1, $2, $3, $4)
RETURNING test_id, owner_id, test_name, description, test_type, created_at
`

type CreateTestParams struct {
	OwnerID     pgtype.UUID
	TestName    string
	Description pgtype.Text
	TestType    TestType
}

// CreateTest
func (q *Queries) CreateTest(ctx context.Context, arg CreateTestParams) (EducationTest, error) {
	row := q.db.QueryRow(ctx, createTest,
		arg.OwnerID,
		arg.TestName,
		arg.Description,
		arg.TestType,
	)
	var i EducationTest
	err := row.Scan(
		&i.TestID,
		&i.OwnerID,
		&i.TestName,
		&i.Description,
		&i.TestType,
		&i.CreatedAt,
	)
	return i, err
}

const getFullTestByOwnerId = `-- name: GetFullTestByOwnerId :many
SELECT t.test_id, owner_id, test_name, description, test_type, t.created_at, q.question_id, q.test_id, question_text, question_type, question_order, q.created_at, answer_id, a.question_id, answer_text, answer_order, score_value, a.created_at
FROM
    education.tests t
JOIN
    education.questions q ON q.test_id = t.test_id
JOIN
    education.answers a ON a.question_id = q.question_id
WHERE t.owner_id = $1
`

type GetFullTestByOwnerIdRow struct {
	TestID        pgtype.UUID
	OwnerID       pgtype.UUID
	TestName      string
	Description   pgtype.Text
	TestType      TestType
	CreatedAt     pgtype.Timestamp
	QuestionID    pgtype.UUID
	TestID_2      pgtype.UUID
	QuestionText  string
	QuestionType  QuestionType
	QuestionOrder int32
	CreatedAt_2   pgtype.Timestamp
	AnswerID      pgtype.UUID
	QuestionID_2  pgtype.UUID
	AnswerText    string
	AnswerOrder   int32
	ScoreValue    pgtype.Numeric
	CreatedAt_3   pgtype.Timestamp
}

// GetFullTestByOwnerId
func (q *Queries) GetFullTestByOwnerId(ctx context.Context, ownerID pgtype.UUID) ([]GetFullTestByOwnerIdRow, error) {
	rows, err := q.db.Query(ctx, getFullTestByOwnerId, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFullTestByOwnerIdRow
	for rows.Next() {
		var i GetFullTestByOwnerIdRow
		if err := rows.Scan(
			&i.TestID,
			&i.OwnerID,
			&i.TestName,
			&i.Description,
			&i.TestType,
			&i.CreatedAt,
			&i.QuestionID,
			&i.TestID_2,
			&i.QuestionText,
			&i.QuestionType,
			&i.QuestionOrder,
			&i.CreatedAt_2,
			&i.AnswerID,
			&i.QuestionID_2,
			&i.AnswerText,
			&i.AnswerOrder,
			&i.ScoreValue,
			&i.CreatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOwner = `-- name: GetOwner :one
SELECT owner_id, owner_name, pass_hash, email, created_at
FROM education.owners
WHERE owner_id = $1
`

// GetOwner
func (q *Queries) GetOwner(ctx context.Context, ownerID pgtype.UUID) (EducationOwner, error) {
	row := q.db.QueryRow(ctx, getOwner, ownerID)
	var i EducationOwner
	err := row.Scan(
		&i.OwnerID,
		&i.OwnerName,
		&i.PassHash,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const getTeacherDisciplinesAndClasses = `-- name: GetTeacherDisciplinesAndClasses :many
SELECT JSON_BUILD_OBJECT(
    'id', sg.discipline_id,
    'title', (SELECT name FROM discipline dis WHERE dis.id = sg.discipline_id),
    'classes', (
        SELECT JSON_AGG(
            JSON_BUILD_OBJECT(
                'id', c.id,
                'title', concat(c.grade, ' ', c.title),
                'main_teacher_id', c.main_teacher_id,
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
WHERE teacher_id = $1
`

// GetTeacherDisciplinesAndClasses
func (q *Queries) GetTeacherDisciplinesAndClasses(ctx context.Context, teacherID pgtype.UUID) ([][]byte, error) {
	rows, err := q.db.Query(ctx, getTeacherDisciplinesAndClasses, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items [][]byte
	for rows.Next() {
		var result []byte
		if err := rows.Scan(&result); err != nil {
			return nil, err
		}
		items = append(items, result)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTestsByOwnerId = `-- name: GetTestsByOwnerId :many
SELECT test_id, owner_id, test_name, description, test_type, created_at
FROM education.tests
WHERE owner_id = $1
`

// GetTestsByOwnerId
func (q *Queries) GetTestsByOwnerId(ctx context.Context, ownerID pgtype.UUID) ([]EducationTest, error) {
	rows, err := q.db.Query(ctx, getTestsByOwnerId, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EducationTest
	for rows.Next() {
		var i EducationTest
		if err := rows.Scan(
			&i.TestID,
			&i.OwnerID,
			&i.TestName,
			&i.Description,
			&i.TestType,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAnswerToTest = `-- name: InsertAnswerToTest :one
INSERT INTO education.answers (question_id, answer_text, answer_order)
VALUES ($1, $2, $3)
RETURNING answer_id, question_id, answer_text, answer_order, score_value, created_at
`

type InsertAnswerToTestParams struct {
	QuestionID  pgtype.UUID
	AnswerText  string
	AnswerOrder int32
}

// InsertAnswerToTest
func (q *Queries) InsertAnswerToTest(ctx context.Context, arg InsertAnswerToTestParams) (EducationAnswer, error) {
	row := q.db.QueryRow(ctx, insertAnswerToTest, arg.QuestionID, arg.AnswerText, arg.AnswerOrder)
	var i EducationAnswer
	err := row.Scan(
		&i.AnswerID,
		&i.QuestionID,
		&i.AnswerText,
		&i.AnswerOrder,
		&i.ScoreValue,
		&i.CreatedAt,
	)
	return i, err
}

const insertQuestionToTest = `-- name: InsertQuestionToTest :one
INSERT INTO education.questions (test_id, question_text, question_order)
VALUES ($1, $2, $3)
RETURNING question_id, test_id, question_text, question_type, question_order, created_at
`

type InsertQuestionToTestParams struct {
	TestID        pgtype.UUID
	QuestionText  string
	QuestionOrder int32
}

// InsertQuestionToTest
func (q *Queries) InsertQuestionToTest(ctx context.Context, arg InsertQuestionToTestParams) (EducationQuestion, error) {
	row := q.db.QueryRow(ctx, insertQuestionToTest, arg.TestID, arg.QuestionText, arg.QuestionOrder)
	var i EducationQuestion
	err := row.Scan(
		&i.QuestionID,
		&i.TestID,
		&i.QuestionText,
		&i.QuestionType,
		&i.QuestionOrder,
		&i.CreatedAt,
	)
	return i, err
}

const listOwner = `-- name: ListOwner :many
SELECT owner_id, owner_name, pass_hash, email, created_at
FROM education.owners
`

// ListOwner
func (q *Queries) ListOwner(ctx context.Context) ([]EducationOwner, error) {
	rows, err := q.db.Query(ctx, listOwner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EducationOwner
	for rows.Next() {
		var i EducationOwner
		if err := rows.Scan(
			&i.OwnerID,
			&i.OwnerName,
			&i.PassHash,
			&i.Email,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
