package service

import (
	"context"

	"walletsvc/internal/entity"
	"walletsvc/pkg/exception"
)

type TransactionService interface {
	GetAllTransactions(ctx context.Context, req GetAllTransactionsReq) (GetAllTransactionsRes, *exception.Exception)
	GetTransaction(ctx context.Context, req GetTransactionReq) (GetTransactionRes, *exception.Exception)
	CreateTransaction(ctx context.Context, req CreateTransactionReq) *exception.Exception
}

type GetAllTransactionsReq struct {
	UserID uint64
}
type GetAllTransactionsRes []*entity.Transaction

type GetTransactionReq struct {
	UserID        uint64
	TransactionID uint64
}
type GetTransactionRes *entity.Transaction

type CreateTransactionReq struct {
	WalletID    uint64  `json:"wallet_id"`
	UserID      uint64  `json:"user_id"`
	ReferenceID string  `json:"reference_id"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
