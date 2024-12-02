package repository

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Transaction struct {
	ID        string    `json:"transactionId" gorm:"primary_key"`
	Amount    string    `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type TransactionRepository interface {
	Insert(ctx context.Context, txn *Transaction) error
}

type transactionRepository struct {
	db        *gorm.DB
	zapLogger *zap.Logger
}

var repo *transactionRepository

func NewTransactionRepository(db *gorm.DB, logger *zap.Logger) *transactionRepository {
	if repo != nil {
		return repo
	}
	repo = &transactionRepository{db: db, zapLogger: logger}
	return repo
}

func (t *transactionRepository) Insert(ctx context.Context, txn *Transaction) error {
	result := t.db.WithContext(ctx).Table("transactions").Create(txn)
	if result.Error != nil {
		t.zapLogger.Fatal("Transaction error: ", zap.Error(result.Error))
		return result.Error
	}
	return nil
}
