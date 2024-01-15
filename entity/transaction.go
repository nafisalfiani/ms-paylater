package entity

const (
	TransactionTypeCredit = "credit"
	TransactionTypeDebit  = "debit"
)

type Transaction struct {
	Id           int     `json:"id" gorm:"primaryKey"`
	UserId       int     `json:"user_id" gorm:"not null"`
	Amount       float64 `json:"amount" gorm:"not null"`
	BalancePrev  float64 `json:"balance_previous"`
	BalanceAfter float64 `json:"balance_after"`
	Type         string  `json:"type" gorm:"not null"`
}

func (t Transaction) CalculateBalance(balance, limit float64) (float64, float64) {
	t.BalancePrev = balance

	if t.Type == TransactionTypeCredit {
		t.BalanceAfter = balance + t.Amount
	} else {
		t.BalanceAfter = balance - t.Amount
	}

	if t.BalanceAfter < 0 {
		t.BalanceAfter = 0
	}

	if t.BalanceAfter > limit {
		t.BalanceAfter = limit
	}

	return t.BalancePrev, t.BalanceAfter
}

type TransactionRequest struct {
	Amount float64 `json:"amount" validate:"required"`
}
