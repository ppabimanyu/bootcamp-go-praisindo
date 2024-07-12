package service

import (
	"context"

	"task_1_crud_user/internal/entity"
	"task_1_crud_user/pkg/exception"
)

type UserService interface {
	GetAllUsers(ctx context.Context) (GetAllUsersRes, *exception.Exception)
	GetDetailUser(ctx context.Context, req GetDetailUserReq) (GetDetailUserRes, *exception.Exception)
	CreateUser(ctx context.Context, req CreateUserReq) *exception.Exception
	UpdateUser(ctx context.Context, req UpdateUserReq) *exception.Exception
	DeleteUser(ctx context.Context, req DeleteUserReq) *exception.Exception
}

type GetDetailUserReq struct {
	UserID uint64 `validate:"required" name:"user_id"`
}

type CreateUserReq struct {
	Name  string `json:"name" validate:"required" name:"name"`
	Email string `json:"email" validate:"required" name:"email"`
}

type UpdateUserReq struct {
	UserID uint64 `validate:"required" name:"user_id"`
	Name   string `json:"name" validate:"required" name:"name"`
	Email  string `json:"email" validate:"required" name:"email"`
}

type DeleteUserReq struct {
	UserID uint64 `validate:"required" name:"user_id"`
}

type GetAllUsersRes []*entity.User

type GetDetailUserRes *entity.User
