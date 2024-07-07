-- name: Createproject :one
INSERT INTO projects (id, title, details)
VALUES ($1, $2, $3)
RETURNING *;