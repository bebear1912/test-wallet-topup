-- +goose Up
-- +goose StatementBegin
INSERT INTO users (username, balance) VALUES
    ('user1', 1000.00),
    ('user2', 500.00),
    ('user3', 2000.00)
ON CONFLICT (username) DO NOTHING;
-- +goose StatementEnd 