package database_sql

import "context"

const createUser = `INSERT INTO users(username, hashed_password, full_name, email)
VALUES($1, $2, $3, $4)
RETURNING username, hashed_password, full_name, email, password_change_at, create_at`

type CreateUserParams struct {
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.UserName, arg.HashedPassword, arg.FullName, arg.Email)
	var i Users
	err := row.Scan(i.UserName, i.HashedPassword, i.FullName, i.Email, i.PasswordChangeAt, i.CreateAt)
	return i, err
}
