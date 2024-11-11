-- +goose Up
CREATE TABLE psychological_type (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE psychological_type;
