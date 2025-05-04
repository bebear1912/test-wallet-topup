-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE username IN ('user1', 'user2', 'user3');
-- +goose StatementEnd 