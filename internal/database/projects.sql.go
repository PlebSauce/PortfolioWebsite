// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: projects.sql

package database

import (
	"context"
)

const createproject = `-- name: Createproject :one
INSERT INTO projects (id, title, details)
VALUES ($1, $2, $3)
RETURNING id, title, details
`

type CreateprojectParams struct {
	ID      string
	Title   string
	Details string
}

func (q *Queries) Createproject(ctx context.Context, arg CreateprojectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createproject, arg.ID, arg.Title, arg.Details)
	var i Project
	err := row.Scan(&i.ID, &i.Title, &i.Details)
	return i, err
}
