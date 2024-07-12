package service

import (
	"context"

	"walletsvc/internal/entity"
	"walletsvc/pkg/exception"
)

type WalletService interface {
	GetAllWallets(ctx context.Context) (GetAllWalletsRes, *exception.Exception)
	GetDetailWallet(ctx context.Context, req GetDetailWalletReq) (GetDetailUserRes, *exception.Exception)
	CreateWallet(ctx context.Context, req CreateWalletReq) *exception.Exception
	UpdateWallet(ctx context.Context, req UpdateWalletReq) *exception.Exception
	DeleteWallet(ctx context.Context, req DeleteWalletReq) *exception.Exception
}

type GetAllWalletsRes []*entity.Wallet

type GetDetailWalletReq struct {
	WalletID uint64 `json:"wallet_id" validate:"required" name:"wallet_id"`
}
type GetDetailUserRes *entity.Wallet

type CreateWalletReq struct {
	UserID uint64 `json:"user_id" validate:"required" name:"user_id"`
}

type UpdateWalletReq struct {
	WalletID uint64  `json:"wallet_id" validate:"required" name:"wallet_id"`
	UserID   uint64  `json:"user_id" validate:"required" name:"user_id"`
	Balance  float64 `json:"balance" validate:"required" name:"balance"`
}

type DeleteWalletReq struct {
	WalletID uint64 `json:"wallet_id" validate:"required" name:"wallet_id"`
}
