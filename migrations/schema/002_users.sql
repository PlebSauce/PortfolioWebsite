-- +goose Up

CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (UUID())
);

-- +goose Down
DROP TABLE users;