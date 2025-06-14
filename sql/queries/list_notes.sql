-- name: ListNotes :many
SELECT id, title, content, created_at
FROM notes
ORDER BY created_at DESC;
