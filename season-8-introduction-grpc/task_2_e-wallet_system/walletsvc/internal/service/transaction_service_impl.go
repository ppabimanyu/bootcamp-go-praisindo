package service

import (
	"context"

	gonanoid "github.com/matoous/go-nanoid"
	"gorm.io/gorm"

	"walletsvc/internal/entity"
	"walletsvc/internal/enums"
	"walletsvc/internal/repository"
	"walletsvc/pkg/exception"
	"walletsvc/pkg/validator"
)

type TransactionServiceImpl struct {
	validator       *validator.Validator
	db              *gorm.DB
	walletRepo      repository.WalletRepository
	transactionRepo repository.TransactionRepository
}

func (s *TransactionServiceImpl) GetAllTransactions(ctx context.Context, req GetAllTransactionsReq) (GetAllTransactionsRes, *exception.Exception) {
	if errs := s.validator.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	transactions, err := s.transactionRepo.FindAll(tx, req.UserID)
	if err != nil {
		return nil, exception.Internal("failed to find transactions", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit transaction", err)
	}

	return transactions, nil
}

func (s *TransactionServiceImpl) GetTransaction(ctx context.Context, req GetTransactionReq) (GetTransactionRes, *exception.Exception) {
	if errs := s.validator.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	transaction, err := s.transactionRepo.FindByID(tx, req.UserID, req.TransactionID)
	if err != nil {
		return nil, exception.Internal("failed to find transactions", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit transaction", err)
	}

	return transaction, nil
}

func (s *TransactionServiceImpl) CreateTransaction(ctx context.Context, req CreateTransactionReq) *exception.Exception {
	if errs := s.validator.Struct(req); errs != nil {
		return exception.InvalidArgument(errs)
	}

	transType := enums.TransactionType(req.Type)
	if !transType.IsValid() {
		return exception.InvalidArgument("invalid transaction type")
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	ref, err := gonanoid.Generate("0123456789", 10)
	if err != nil {
		return exception.Internal("failed to generate reference id", err)
	}

	if err := s.transactionRepo.Create(tx, &entity.Transaction{
		WalletID:    req.WalletID,
		UserID:      req.UserID,
		ReferenceID: ref,
		Type:        transType.String(),
		Amount:      req.Amount,
		Status:      "success",
		Description: req.Description,
	}); err != nil {
		return exception.Internal("failed to create transactions", err)
	}

	wallet, err := s.walletRepo.FindByID(tx, req.WalletID)
	if err != nil {
		return exception.Internal("failed to find wallet", err)
	}
	if wallet == nil {
		return exception.NotFound("wallet not found")
	}

	if transType == enums.TransactionTypeDeposit {
		wallet.Balance += req.Amount
	} else {
		wallet.Balance -= req.Amount
	}

	if err := s.walletRepo.Update(tx, wallet); err != nil {
		return exception.Internal("failed to update wallet", err)
	}

	if err := tx.Commit().Error; err != nil {
		return exception.Internal("failed to commit transaction", err)
	}

	return nil
}
