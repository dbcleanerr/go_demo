package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"go_demo/simplebank/util"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.ID, account1.ID)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, account1.Balance)
	require.Equal(t, account2.Currency, account1.Currency)
	require.Equal(t, account2.CreatedAt, account1.CreatedAt)

	require.WithinDuration(t, account2.CreatedAt, account1.CreatedAt, 0)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.ID, account1.ID)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, arg.Balance)
	require.Equal(t, account2.Currency, account1.Currency)
	require.Equal(t, account2.CreatedAt, account1.CreatedAt)

	require.WithinDuration(t, account2.CreatedAt, account1.CreatedAt, 0)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	account3, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account3)

}
