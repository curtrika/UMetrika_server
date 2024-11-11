-- +goose Up
CREATE TABLE IF NOT EXISTS psychological_type (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE psychological_type;
