package entity

import "time"

const (
	_transactionsTableName = "transactions"
)

type Transaction struct {
	ID          uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	WalletID    uint64     `gorm:"column:wallet_id;"`
	UserID      uint64     `gorm:"column:user_id;"`
	ReferenceID string     `gorm:"column:reference_id;"`
	Type        string     `gorm:"column:type;check:type in ('deposit','withdraw')"`
	Amount      float64    `gorm:"column:amount;"`
	Status      string     `gorm:"column:status;check:status in ('pending','success','failed')"`
	Description string     `gorm:"column:description;"`
	CreatedAt   *time.Time `gorm:"column:created_at;"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;"`
	Wallet      Wallet     `gorm:"foreignKey:WalletID;references:ID"`
}

func (w *Transaction) TableName() string {
	return _transactionsTableName
}
