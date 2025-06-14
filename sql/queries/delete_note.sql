-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;
