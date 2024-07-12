package service

import (
	"context"

	"gorm.io/gorm"

	"usersvc/internal/entity"
	"usersvc/internal/enums"
	"usersvc/internal/repository"
	"usersvc/pkg/exception"
	"usersvc/pkg/validator"
)

type UserServiceImpl struct {
	validator *validator.Validator
	db        *gorm.DB
	userRepo  repository.UserRepository
}

func NewUserServiceImpl(
	validator *validator.Validator,
	db *gorm.DB,
	userRepo repository.UserRepository,
) UserService {
	return &UserServiceImpl{
		validator: validator,
		db:        db,
		userRepo:  userRepo,
	}
}

func (s *UserServiceImpl) GetAllUsers(ctx context.Context) (GetAllUsersRes, *exception.Exception) {
	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	users, err := s.userRepo.FindAll(tx)
	if err != nil {
		return nil, exception.Internal("failed to get all users", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit transaction", err)
	}

	return users, nil
}

func (s *UserServiceImpl) GetDetailUser(ctx context.Context, req GetDetailUserReq) (GetDetailUserRes, *exception.Exception) {
	if errs := s.validator.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := s.userRepo.FindByID(tx, req.UserID)
	if err != nil {
		return nil, exception.Internal("failed to get detail user", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit transaction", err)
	}

	return user, nil
}

func (s *UserServiceImpl) CreateCustomer(ctx context.Context, req CreateCustomerReq) *exception.Exception {
	if errs := s.validator.Struct(req); errs != nil {
		return exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.userRepo.Create(tx, &entity.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
		Type:  string(enums.Customer),
	}); err != nil {
		return exception.Internal("failed to create customer", err)
	}

	if err := tx.Commit().Error; err != nil {
		return exception.Internal("failed to commit transaction", err)
	}

	return nil
}

func (s *UserServiceImpl) CreateMerchant(ctx context.Context, req CreateMerchantReq) *exception.Exception {
	if errs := s.validator.Struct(req); errs != nil {
		return exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.userRepo.Create(tx, &entity.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
		Type:  string(enums.Merchant),
	}); err != nil {
		return exception.Internal("failed to create merchant", err)
	}

	if err := tx.Commit().Error; err != nil {
		return exception.Internal("failed to commit transaction", err)
	}

	return nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, req UpdateUserReq) *exception.Exception {
	if errs := s.validator.Struct(req); errs != nil {
		return exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := s.userRepo.FindByID(tx, req.UserID)
	if err != nil {
		return exception.Internal("failed to get detail user", err)
	}
	if user == nil {
		return exception.NotFound("user not found")
	}

	if err := s.userRepo.Update(tx, &entity.User{
		Name:  req.Name,
		Email: req.Email,
	}); err != nil {
		return exception.Internal("failed to update user", err)
	}

	if err := tx.Commit().Error; err != nil {
		return exception.Internal("failed to commit transaction", err)
	}

	return nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, req DeleteUserReq) *exception.Exception {
	if errs := s.validator.Struct(req); errs != nil {
		return exception.InvalidArgument(errs)
	}

	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := s.userRepo.FindByID(tx, req.UserID)
	if err != nil {
		return exception.Internal("failed to get detail user", err)
	}
	if user == nil {
		return exception.NotFound("user not found")
	}

	if err := s.userRepo.Delete(tx, user); err != nil {
		return exception.Internal("failed to delete user", err)
	}

	if err := tx.Commit().Error; err != nil {
		return exception.Internal("failed to commit transaction", err)
	}

	return nil
}
