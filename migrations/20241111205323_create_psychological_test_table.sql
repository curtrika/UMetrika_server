-- +goose Up
CREATE TABLE psychological_test (
    id SERIAL PRIMARY KEY,
    first_question_id INT REFERENCES question(id) ON DELETE SET NULL,
    type_id INT REFERENCES psychological_type(id) ON DELETE SET NULL,
    owner_id INT NOT NULL,
    title VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE psychological_test;
