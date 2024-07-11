package repository

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"

	"intro-grpc-task/internal/entity"
)

const (
	failedFindUser   = "failed find user"
	failedCreateUser = "failed create user"
	failedUpdateUser = "failed update user"
	failedDeleteUser = "failed delete user"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindAll(tx *gorm.DB) ([]*entity.User, error) {
	var users []*entity.User
	if err := tx.Find(&users).Error; err != nil {
		slog.Error(failedFindUser, err)
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) FindByID(tx *gorm.DB, id uint64) (*entity.User, error) {
	var user entity.User
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		slog.Error(failedFindUser, err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(tx *gorm.DB, user *entity.User) error {
	if err := tx.Create(user).Error; err != nil {
		slog.Error(failedCreateUser, err)
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Update(tx *gorm.DB, user *entity.User) error {
	if err := tx.Select("*").Updates(user).Error; err != nil {
		slog.Error(failedUpdateUser, err)
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Delete(tx *gorm.DB, user *entity.User) error {
	if err := tx.Delete(user).Error; err != nil {
		slog.Error(failedDeleteUser, err)
		return err
	}
	return nil
}
