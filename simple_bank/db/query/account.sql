-- name: CreateAccount :one
   INSERT INTO accounts(owner, balance, currency)
   VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteAccount :exec
DELETE
  FROM accounts
 WHERE id = $1;

-- name: UpdateAccount :one
   UPDATE accounts
      SET balance = $2
    WHERE id = $1
RETURNING *;

-- name: GetAccount :one
SELECT *
  FROM accounts
 WHERE id = $1;

-- name: ListAccount :many
SELECT *
  FROM accounts
OFFSET $1 LIMIT $2;