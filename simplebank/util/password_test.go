package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	password1 := RandomString(8)
	hashPassword1, err := HashPassword(password1)

	require.NoError(t, err)
	require.NotEmpty(t, hashPassword1)

	err = CheckPassword(password1, hashPassword1)
	require.NoError(t, err)

	// 错误的密码
	wrongPassword := RandomString(8)
	err = CheckPassword(wrongPassword, hashPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	// 同一个密码生成不同的 hash password
	hashPassword2, err := HashPassword(password1)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword2)

	err = CheckPassword(password1, hashPassword2)
	require.NoError(t, err)

	require.NotEqual(t, hashPassword1, hashPassword2)
}
