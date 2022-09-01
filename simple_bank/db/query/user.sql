-- name: CreateUser :one
   INSERT INTO users(username, hashed_password, full_name, email)
   VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteUser :one
   DELETE
     FROM users
    WHERE username = $1
RETURNING *;

-- name: UpdateUser :one
   UPDATE users
      SET hashed_password     = $1,
          full_name           = $2,
          email               = $3,
          password_changed_at = NOW()
    WHERE username = $4
RETURNING *;

-- name: GetUser :one
SELECT *
  FROM users
 WHERE username = $1;

-- name: ListUser :many
SELECT *
  FROM users
OFFSET $1 LIMIT $2;