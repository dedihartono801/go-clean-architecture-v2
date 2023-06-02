package repository

import (
	"github.com/dedihartono801/go-clean-architecture-v2/internal/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *gorm.DB, transaction *entity.Transaction) error
}

type transactionRepository struct {
	database *gorm.DB
}

func NewTransactionRepository(database *gorm.DB) TransactionRepository {
	return &transactionRepository{database}
}

func (r *transactionRepository) Create(tx *gorm.DB, transaction *entity.Transaction) error {
	return tx.Table("transaction").Create(transaction).Error
}
