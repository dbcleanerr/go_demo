-- name: CreateTransfer :one
   INSERT INTO transfers(id, from_account_id, to_account_id, amount)
   VALUES ($1, $2, $3, $3)
RETURNING *;

-- name: GetTransfer :one
SELECT *
  FROM transfers
 WHERE id = $1;

-- name: ListTransfer :many
SELECT *
  FROM transfers
OFFSET $1 LIMIT $2;