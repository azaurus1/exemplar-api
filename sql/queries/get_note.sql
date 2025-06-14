-- name: GetNote :one
SELECT id, title, content, created_at
FROM notes
WHERE id = $1;
