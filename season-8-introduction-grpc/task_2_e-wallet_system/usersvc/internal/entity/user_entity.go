package entity

import "time"

const (
	_usersTableName = "wallet"
)

type User struct {
	ID        uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string     `gorm:"column:name;not null"`
	Email     string     `gorm:"column:email;not null; unique"`
	Phone     string     `gorm:"column:phone"`
	Type      string     `gorm:"column:type;check:type IN ('customer', 'merchant')"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (u *User) TableName() string {
	return _usersTableName
}
