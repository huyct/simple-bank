package store

import (
	"context"
	"testing"

	"github.com/duckhue01/back-end/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)

}

func TestAddAccountBalance(t *testing.T) {
	initAcc := createRandomAccount(t)
	UpdatedAcc, err := testQueries.AddAccountBalance(context.Background(), AddAccountBalanceParams{
		ID:     initAcc.ID,
		Amount: 10,
	})

	require.NoError(t, err)
	require.Equal(t, initAcc.Balance+10, UpdatedAcc.Balance)

}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)
}

func TestGetAccount(t *testing.T) {
	createdAcc := createRandomAccount(t)
	gotAcc, err := testQueries.GetAccount(context.Background(), createdAcc.ID)
	require.NoError(t, err)

	require.Equal(t, createdAcc, gotAcc)
}

func TestListAccounts(t *testing.T) {
	var firstOwner string
	for i := 0; i < 10; i++ {

		if i == 0 {
			firstOwner = createRandomAccount(t).Owner
		}
	}

	arg := ListAccountsParams{
		Limit:  1,
		Offset: 0,
		Owner:  firstOwner,
	}
	accs, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accs, 1)

}

func TestUpdateAccount(t *testing.T) {
	createdAcc := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAcc.ID,
		Balance: 10,
	}
	updatedAcc, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, updatedAcc.Balance, arg.Balance)
}

func TestAccountForUpdate(t *testing.T) {

	createdAcc := createRandomAccount(t)
	gotAcc, err := testQueries.GetAccountForUpdate(context.Background(), createdAcc.ID)

	require.NoError(t, err)
	require.Equal(t, createdAcc, gotAcc)
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	return account

}
