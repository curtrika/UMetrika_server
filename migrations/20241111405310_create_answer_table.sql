-- +goose Up
CREATE TABLE IF NOT EXISTS answer (
    id SERIAL PRIMARY KEY,
    next_answer_id INT REFERENCES answer(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE answer;
