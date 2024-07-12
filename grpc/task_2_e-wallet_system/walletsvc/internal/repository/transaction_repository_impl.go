package repository

import (
	"errors"

	"gorm.io/gorm"

	"walletsvc/internal/entity"
	"walletsvc/pkg/logger"
)

type TransactionRepositoryImpl struct {
	log logger.Logger
}

func NewTransactionRepositoryImpl(log logger.Logger) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{log: log}
}

func (r *TransactionRepositoryImpl) FindAll(tx *gorm.DB, userID uint64) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	if err := tx.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {

		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepositoryImpl) FindByID(tx *gorm.DB, userID, transactionID uint64) (*entity.Transaction, error) {
	var transaction entity.Transaction
	if err := tx.Where("user_id = ? AND id = ?", userID, transactionID).First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		r.log.Error("failed to find transaction", "error", err)
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepositoryImpl) Create(tx *gorm.DB, transaction *entity.Transaction) error {
	if err := tx.Create(transaction).Error; err != nil {
		r.log.Error("failed to create transaction", "error", err)
		return err
	}
	return nil
}
