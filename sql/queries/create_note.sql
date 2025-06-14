-- name: CreateNote :one
INSERT INTO notes (title, content)
VALUES ($1, $2)
RETURNING id, title, content, created_at;
