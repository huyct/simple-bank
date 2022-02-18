package store

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	acc01 := createRandomAccount(t)
	acc02 := createRandomAccount(t)

	amount := int64(10)
	n := 2

	errChan := make(chan error)
	resultChan := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			re, err := store.TransferTx(context.WithValue(context.Background(), txKey, txName), TransferTxParams{
				FromAccountID: acc01.ID,
				ToAccountID:   acc02.ID,
				Amount:        amount,
			})

			errChan <- err
			resultChan <- re
		}()
	}

	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errChan
		require.NoError(t, err)

		re := <-resultChan
		require.NotEmpty(t, re)

		// check transfer
		require.NotEmpty(t, re.Transfer)

		require.Equal(t, acc01.ID, re.Transfer.FromAccountID)
		require.Equal(t, acc02.ID, re.Transfer.ToAccountID)
		require.Equal(t, amount, re.Transfer.Amount)

		// check entries
		require.NotEmpty(t, re.FromEntry)
		require.NotEmpty(t, re.ToEntry)

		require.Equal(t, acc01.ID, re.FromEntry.AccountID)
		require.Equal(t, acc02.ID, re.ToEntry.AccountID)

		require.Equal(t, amount, re.ToEntry.Amount)
		require.Equal(t, -amount, re.FromEntry.Amount)

		// check accounts
		require.NotEmpty(t, re.FromAccount)
		require.NotEmpty(t, re.ToAccount)

		require.Equal(t, acc01.ID, re.FromAccount.ID)
		require.Equal(t, acc02.ID, re.ToAccount.ID)

		diff01 := acc01.Balance - re.FromAccount.Balance
		diff02 := re.ToAccount.Balance - acc02.Balance

		require.Equal(t, diff01, diff02)
		require.True(t, diff01%amount == 0)

		k := int(diff01 / amount)

		require.True(t, k <= n && k >= 1)

		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedAcc01, err := store.GetAccount(context.Background(), acc01.ID)
	require.NoError(t, err)

	updatedAcc02, err := store.GetAccount(context.Background(), acc02.ID)
	require.NoError(t, err)

	require.Equal(t, acc01.Balance, updatedAcc01.Balance+(amount*int64(n)))
	require.Equal(t, acc02.Balance, updatedAcc02.Balance-(amount*int64(n)))

}
