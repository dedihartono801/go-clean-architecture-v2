package repository

import (
	"gorm.io/gorm"
)

type DbTransactionRepository interface {
	BeginTransaction() (*gorm.DB, error)
	CommitTransaction(tx *gorm.DB) error
}

type dbTransactionRepository struct {
	database *gorm.DB
}

func NewDbTransactionRepository(database *gorm.DB) DbTransactionRepository {
	return &dbTransactionRepository{database}
}

// BeginTransaction starts a new transaction and returns the begin operation
func (r *dbTransactionRepository) BeginTransaction() (*gorm.DB, error) {
	tx := r.database.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

// CommitTransaction commits a transaction and returns the commit operation
func (r *dbTransactionRepository) CommitTransaction(tx *gorm.DB) error {
	err := tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}
