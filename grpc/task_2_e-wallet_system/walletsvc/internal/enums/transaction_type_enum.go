package enums

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeWithdrawal TransactionType = "withdrawal"
)

func (t TransactionType) IsValid() bool {
	return t == TransactionTypeDeposit || t == TransactionTypeWithdrawal
}

func (t TransactionType) String() string {
	return string(t)
}
