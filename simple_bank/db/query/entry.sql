-- name: CreateEntry :one
   INSERT INTO entries(id, account_id, amount)
   VALUES ($1, $2, $3)
RETURNING *;

-- name: GetEntry :one
SELECT *
  FROM entries
 WHERE id = $1;

-- name: ListEntry :many
SELECT *
  FROM entries
OFFSET $1 LIMIT $2;