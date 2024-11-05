-- +goose Up
-- +goose StatementBegin
INSERT INTO apps (id, name, secret)
VALUES (1, 'test', 'test-secret')
ON CONFLICT DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table apps;
-- +goose StatementEnd
