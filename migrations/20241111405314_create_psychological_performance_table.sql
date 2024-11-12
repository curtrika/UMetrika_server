-- +goose Up
CREATE TABLE IF NOT EXISTS psychological_performance (
    id SERIAL PRIMARY KEY,
    owner_id INT NOT NULL,
    psychological_test_id INT REFERENCES psychological_test(id) ON DELETE CASCADE,
    started_at TIMESTAMPTZ NOT NULL
);

-- +goose Down
DROP TABLE psychological_performance;
