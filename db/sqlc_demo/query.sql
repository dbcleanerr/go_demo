-- name: GetAuthor :one
SELECT *
  FROM authors
 WHERE id = $1 LIMIT 1;

-- name: ListAuthorByIDs :many
SELECT *
  FROM authors
 WHERE id = ANY($1::int[]);

-- name: ListAuthors :many
SELECT * FROM authors
 ORDER BY name;

-- name: CreateGetAuthor :one
INSERT INTO authors (name, bio) VALUES ($1, $2)
RETURNING *;

-- name: CreateAuthor :exec
INSERT INTO authors (name, bio) VALUES ($1, $2);

-- name: DeleteAuthor :exec
DELETE FROM authors
 WHERE id = $1;

-- name: DeleteGetAuthor :one
DELETE FROM authors
 WHERE id = $1
RETURNING *;

-- name: UpdateGetAuthor :one
UPDATE authors
   SET name = $2,
       bio = $3
 WHERE id = $1
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors
   SET name = $2,
       bio = $3
 WHERE id = $1;