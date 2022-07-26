package db

import "context"

const createUser = `-- name: CreateUser :one
   INSERT INTO users(username, hashed_password, full_name, email)
   VALUES ($1, $2, $3, $4)
RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.HashedPassword, arg.FullName, arg.Email)
	var i User
	err := row.Scan(
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangeAt,
		&i.CreateAt)

	return i, err
}
