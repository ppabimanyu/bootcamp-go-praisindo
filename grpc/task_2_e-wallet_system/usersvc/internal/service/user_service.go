package service

import (
	"context"

	"usersvc/internal/entity"
	"usersvc/pkg/exception"
)

type UserService interface {
	GetAllUsers(ctx context.Context) (GetAllUsersRes, *exception.Exception)
	GetDetailUser(ctx context.Context, req GetDetailUserReq) (GetDetailUserRes, *exception.Exception)
	CreateCustomer(ctx context.Context, req CreateCustomerReq) *exception.Exception
	CreateMerchant(ctx context.Context, req CreateMerchantReq) *exception.Exception
	UpdateUser(ctx context.Context, req UpdateUserReq) *exception.Exception
	DeleteUser(ctx context.Context, req DeleteUserReq) *exception.Exception
}

type GetDetailUserReq struct {
	UserID uint64 `validate:"required" name:"user_id"`
}

type CreateCustomerReq struct {
	Name  string `json:"name" validate:"required" name:"name"`
	Email string `json:"email" validate:"required" name:"email"`
	Phone string `json:"phone" validate:"required" name:"phone"`
}

type CreateMerchantReq struct {
	Name  string `json:"name" validate:"required" name:"name"`
	Email string `json:"email" validate:"required" name:"email"`
	Phone string `json:"phone" validate:"required" name:"phone"`
}

type UpdateUserReq struct {
	UserID uint64 `validate:"required" name:"user_id"`
	Name   string `json:"name" validate:"required" name:"name"`
	Email  string `json:"email" validate:"required" name:"email"`
	Phone  string `json:"phone" validate:"required" name:"phone"`
	Type   string `json:"type" validate:"required" name:"type"`
}

type DeleteUserReq struct {
	UserID uint64 `validate:"required" name:"user_id"`
}

type GetAllUsersRes []*entity.User

type GetDetailUserRes *entity.User
