package repository

import (
	"errors"

	"gorm.io/gorm"

	"walletsvc/internal/entity"
	"walletsvc/pkg/logger"
)

const (
	failedFindWallet   = "failed find wallet"
	failedCreateWallet = "failed create wallet"
	failedUpdateWallet = "failed update wallet"
	failedDeleteWallet = "failed delete wallet"
)

type WalletRepositoryImpl struct {
	log logger.Logger
}

func NewWalletRepositoryImpl(log logger.Logger) WalletRepository {
	return &WalletRepositoryImpl{
		log: log,
	}
}

func (r *WalletRepositoryImpl) FindAll(tx *gorm.DB) ([]*entity.Wallet, error) {
	var wallets []*entity.Wallet
	if err := tx.Find(&wallets).Error; err != nil {
		r.log.Error(failedFindWallet, err)
		return nil, err
	}
	return wallets, nil
}

func (r *WalletRepositoryImpl) FindByID(tx *gorm.DB, id uint64) (*entity.Wallet, error) {
	var wallet entity.Wallet
	if err := tx.Where("id = ?", id).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		r.log.Error(failedFindWallet, err)
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepositoryImpl) Create(tx *gorm.DB, wallet *entity.Wallet) error {
	if err := tx.Create(wallet).Error; err != nil {
		r.log.Error(failedCreateWallet, err)
		return err
	}
	return nil
}

func (r *WalletRepositoryImpl) Update(tx *gorm.DB, wallet *entity.Wallet) error {
	if err := tx.Select("*").Updates(wallet).Error; err != nil {
		r.log.Error(failedUpdateWallet, err)
		return err
	}
	return nil
}

func (r *WalletRepositoryImpl) Delete(tx *gorm.DB, wallet *entity.Wallet) error {
	if err := tx.Delete(wallet).Error; err != nil {
		r.log.Error(failedDeleteWallet, err)
		return err
	}
	return nil
}
