-- name: UpdateNote :one
UPDATE notes
SET title = $2,
    content = $3
WHERE id = $1
RETURNING id, title, content, created_at;
