-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title;

-- name: CreateBook :one
INSERT INTO books (
          title, status
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;
