package repository

import (
	"gorm.io/gorm"

	"task_1_crud_user/internal/entity"
)

type UserRepository interface {
	FindAll(tx *gorm.DB) ([]*entity.User, error)
	FindByID(tx *gorm.DB, userID uint64) (*entity.User, error)
	Create(tx *gorm.DB, user *entity.User) error
	Update(tx *gorm.DB, user *entity.User) error
	Delete(tx *gorm.DB, user *entity.User) error
}
