-- +goose Up

CREATE TABLE projects (
    id CHAR(36) PRIMARY KEY,
    title TEXT NOT NULL,
    details TEXT NOT NULL
);

-- +goose Down
DROP TABLE projects;