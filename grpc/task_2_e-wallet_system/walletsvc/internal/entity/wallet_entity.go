package entity

import "time"

const (
	_walletsTableName = "wallet"
)

type Wallet struct {
	ID        uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    uint64     `gorm:"column:user_id;not null"`
	Balance   float64    `gorm:"column:balance"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (w *Wallet) TableName() string {
	return _walletsTableName
}
