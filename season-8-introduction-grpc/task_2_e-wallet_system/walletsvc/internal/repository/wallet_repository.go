package repository

import (
	"gorm.io/gorm"

	"walletsvc/internal/entity"
)

type WalletRepository interface {
	FindAll(tx *gorm.DB) ([]*entity.Wallet, error)
	FindByID(tx *gorm.DB, walletID uint64) (*entity.Wallet, error)
	Create(tx *gorm.DB, wallet *entity.Wallet) error
	Update(tx *gorm.DB, wallet *entity.Wallet) error
	Delete(tx *gorm.DB, wallet *entity.Wallet) error
}
