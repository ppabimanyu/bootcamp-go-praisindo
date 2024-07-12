package repository

import (
	"gorm.io/gorm"

	"walletsvc/internal/entity"
)

type TransactionRepository interface {
	FindAll(tx *gorm.DB, userID uint64) ([]*entity.Transaction, error)
	FindByID(tx *gorm.DB, userID, transactionID uint64) (*entity.Transaction, error)
	Create(tx *gorm.DB, transaction *entity.Transaction) error
}
