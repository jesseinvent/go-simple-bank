package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// Creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %c, rb err %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

// Contains the input parameters of he transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Performs a money transfer from one account to the other
// It creates a transfer record, add account entries, and update accounts balance within a single database transaction
// func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {}
