package service

import (
	"context"

	"gorm.io/gorm"

	"walletsvc/internal/entity"
	"walletsvc/internal/repository"
	"walletsvc/pkg/exception"
	"walletsvc/pkg/validator"
)

type UserServiceImpl struct {
	validator  *validator.Validator
	db         *gorm.DB
	walletRepo repository.WalletRepository
}

func NewUserServiceImpl(
	validator *validator.Validator,
	db *gorm.DB,
	walletRepo repository.WalletRepository,
) *UserServiceImpl {
	return &UserServiceImpl{
		validator:  validator,
		db:         db,
		walletRepo: walletRepo,
	}
}

func (s *UserServiceImpl) GetAllWallets(ctx context.Context) (GetAllWalletsRes, *exception.Exception) {
	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	wallets, err := s.walletRepo.FindAll(tx)
	if err != nil {
		return nil, exception.Internal("failed to get all wallet", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit transaction", err)
	}

	return wallets, nil
}

func (s *UserServiceImpl) GetDetailWallet(ctx context.Context, req GetDetailWalletReq) (GetDetailUserRes, *exception.Exception) {
	if errs := s.validator.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	wallet, err := s.walletRepo.FindByID(tx, req.WalletID)
	if err != nil {
		return nil, exception.Internal("failed to get detail wallet", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit transaction", err)
	}

	return wallet, nil
}

func (s *UserServiceImpl) CreateWallet(ctx context.Context, req CreateWalletReq) *exception.Exception {
	if errs := s.validator.Struct(req); errs != nil {
		return exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.walletRepo.Create(tx, &entity.Wallet{
		UserID:  req.UserID,
		Balance: 0,
	}); err != nil {
		return exception.Internal("failed to create wallet", err)
	}

	if err := tx.Commit().Error; err != nil {
		return exception.Internal("failed to commit transaction", err)
	}

	return nil
}

func (s *UserServiceImpl) UpdateWallet(ctx context.Context, req UpdateWalletReq) *exception.Exception {
	if errs := s.validator.Struct(req); errs != nil {
		return exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	wallet, err := s.walletRepo.FindByID(tx, req.WalletID)
	if err != nil {
		return exception.Internal("failed to get detail wallet", err)
	}
	if wallet == nil {
		return exception.NotFound("wallet not found")
	}

	if err := s.walletRepo.Update(tx, &entity.Wallet{
		UserID:  req.UserID,
		Balance: req.Balance,
	}); err != nil {
		return exception.Internal("failed to update wallet", err)
	}

	if err := tx.Commit().Error; err != nil {
		return exception.Internal("failed to commit transaction", err)
	}

	return nil
}

func (s *UserServiceImpl) DeleteWallet(ctx context.Context, req DeleteWalletReq) *exception.Exception {
	if errs := s.validator.Struct(req); errs != nil {
		return exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	wallet, err := s.walletRepo.FindByID(tx, req.WalletID)
	if err != nil {
		return exception.Internal("failed to get detail wallet", err)
	}
	if wallet == nil {
		return exception.NotFound("wallet not found")
	}

	if err := s.walletRepo.Delete(tx, wallet); err != nil {
		return exception.Internal("failed to delete wallet", err)
	}

	if err := tx.Commit().Error; err != nil {
		return exception.Internal("failed to commit transaction", err)
	}

	return nil
}
