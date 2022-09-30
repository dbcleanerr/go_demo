package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:       "1001",
		HashedPassword: "welcome",
		FullName:       "1001-1002",
		Email:          "1001@email.com",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
