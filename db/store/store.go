package store

import (
	"database/sql"
	"learn/back-end/db/sqlc"
)

type Store struct {
	*sqlc.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: sqlc.New(db),
	}
}

// type TransferTxParams struct {
// 	FromAccountID int64 `json:"from_account_id"`
// 	ToAccountID   int64 `json:"to_account_id"`
// 	Amount        int64 `json:"amount"`
// }

// type TransferTxResult struct {
// 	Transfer    sqlc.Transfer `json:"transfer"`
// 	FromAccount sqlc.Account  `json:"from_account"`
// 	ToAccount   sqlc.Account  `json:"to_account"`
// 	FromEntry   sqlc.Entry    `json:"from_entry"`
// 	ToEntry     sqlc.Entry    `json:"to_entry"`
// }
