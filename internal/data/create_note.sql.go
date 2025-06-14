// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: create_note.sql

package data

import (
	"context"
	"database/sql"
)

const createNote = `-- name: CreateNote :one
INSERT INTO notes (title, content)
VALUES ($1, $2)
RETURNING id, title, content, created_at
`

type CreateNoteParams struct {
	Title   string         `json:"title"`
	Content sql.NullString `json:"content"`
}

func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, createNote, arg.Title, arg.Content)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}
