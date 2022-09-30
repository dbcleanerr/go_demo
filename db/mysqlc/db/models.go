package db

import "time"

type User struct {
	UserName         string    `json:"user_name"`
	HashedPassword   string    `json:"hashed_password"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreateAt         time.Time `json:"create_at"`
}
