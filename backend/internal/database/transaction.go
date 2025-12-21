package database

import (
	"context"
	"database/sql"
)

// TransactionManagerImpl implements TransactionManager
type TransactionManagerImpl struct {
	db *sql.DB
}

// NewTransactionManager creates a new TransactionManagerImpl
func NewTransactionManager(db *sql.DB) *TransactionManagerImpl {
	return &TransactionManagerImpl{db: db}
}

// BeginTx starts a new transaction
func (tm *TransactionManagerImpl) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return tm.db.BeginTx(ctx, nil)
}

// Commit commits the transaction
func (tm *TransactionManagerImpl) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

// Rollback rolls back the transaction
func (tm *TransactionManagerImpl) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
