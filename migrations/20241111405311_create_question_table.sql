-- +goose Up
CREATE TABLE IF NOT EXISTS question (
    id SERIAL PRIMARY KEY,
    next_question_id INT REFERENCES question(id) ON DELETE SET NULL,
    number INT NOT NULL,
    first_answer_id INT REFERENCES answer(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE question;
