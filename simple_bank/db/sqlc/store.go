package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provide all function to execute db query and transaction
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore create a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// ExecTx execute a function within a database transaction
func (s *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// New 可以 db 和 tx
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err:%s, rb err:%s", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
