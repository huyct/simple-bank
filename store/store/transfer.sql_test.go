package store

import (
	"context"
	"testing"

	"github.com/duckhue01/back-end/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test create transfer"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc1 := createRandomAccount(t)
			acc2 := createRandomAccount(t)
			createRandomTransfer(t, acc1, acc2)
		})
	}

}

func TestGetTransfer(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "text get transfer", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc1 := createRandomAccount(t)
			acc2 := createRandomAccount(t)
			expectedTransfer := createRandomTransfer(t, acc1, acc2)
			gotTransfer, err := testQueries.GetTransfer(context.Background(), expectedTransfer.ID)

			if (err != nil) != tt.wantErr {
				t.Errorf("testQueries.GetTransfer() error = %v, wantError %v", err, tt.wantErr)
				return
			}
			require.Equal(t, expectedTransfer, gotTransfer)

		})
	}

}

func TestListTransfers(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
		err     error
	}{
		{
			name: "test list transfer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc1 := createRandomAccount(t)
			acc2 := createRandomAccount(t)
			expectedTransfer := createRandomTransfer(t, acc1, acc2)
			gotTransfer, err := testQueries.ListTransfers(context.Background(), ListTransfersParams{
				FromAccountID: acc1.ID,
				ToAccountID:   acc2.ID,
				Limit:         1,
				Offset:        0,
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("testQueries.Listtransfers() error = %v, wantError %v", err, tt.wantErr)
				return
			}

			require.Equal(t, tt.err, err)
			require.Len(t, gotTransfer, 1)
			require.Equal(t, expectedTransfer, gotTransfer[0])

		})
	}
}

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}
