package store

import (
	"context"
	"testing"

	"github.com/duckhue01/back-end/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
		err     error
	}{
		{
			"test create entry",
			false,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			createRandomEntry(t, tt.wantErr, tt.err)

		})
	}

}

func TestGetEntry(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
		err     error
	}{
		{
			"test get entry test case 1",
			false,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedEntry := createRandomEntry(t, false, nil)
			gotEntry, err := testQueries.GetEntry(context.Background(), expectedEntry.ID)

			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.ListTransfers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.Equal(t, expectedEntry, gotEntry)

		})
	}

}

func TestListEntry(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
		err     error
	}{
		{
			"test list entry",
			false,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry := createRandomEntry(t, false, nil)

			entries, err := testQueries.ListEntries(context.Background(), ListEntriesParams{
				AccountID: entry.AccountID,
				Limit:     1,
				Offset:    0,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("testQueries.ListEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.err, err)
			require.Len(t, entries, 1)

		})
	}

}

func createRandomEntry(t *testing.T, wantErr bool, expectedErr error) Entry {
	acc := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	if wantErr {
		t.Errorf("Queries.ListTransfers() error = %v, wantErr %v", err, wantErr)
	} else {
		require.NoError(t, err)
		require.NotEmpty(t, entry)

		require.Equal(t, arg.AccountID, entry.AccountID)
		require.Equal(t, arg.Amount, entry.Amount)

		require.NotZero(t, entry.ID)
		require.NotZero(t, entry.CreatedAt)
	}

	return entry
}
