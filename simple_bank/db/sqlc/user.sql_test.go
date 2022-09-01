package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simple_bank/util"
	"testing"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "welcome",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashedPassword, arg.HashedPassword)
	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Email, arg.Email)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.PasswordChangedAt)

	return user

}

func TestQueries_CreateUser(t *testing.T) {
	createRandomUser(t)
}
