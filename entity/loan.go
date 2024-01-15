package entity

import "strings"

type Loan struct {
	Id      int     `json:"id" gorm:"primaryKey"`
	UserId  int     `json:"user_id" gorm:"foreignKey"`
	Balance float64 `json:"balance"`
	Limit   float64 `json:"limit"`
	Tier    string  `json:"tier"`
}

func (l Loan) CalculateBalance(transactionType string, amount float64) float64 {
	if transactionType == TransactionTypeCredit {
		l.Balance = l.Balance + amount
	} else {
		l.Balance = l.Balance - amount
	}

	if l.Balance < 0 {
		l.Balance = 0
	}

	if l.Balance > l.Limit {
		l.Balance = l.Limit
	}

	return l.Balance
}

type LoanRequest struct {
	Tier string `json:"tier" validate:"required"`
}

func (l LoanRequest) AssignTierAndLimit() (string, float64) {
	var tier string
	var limit float64
	switch strings.ToLower(l.Tier) {
	case "bronze":
		tier = l.Tier
		limit = 5000000
	case "silver":
		tier = l.Tier
		limit = 10000000
	case "gold":
		tier = l.Tier
		limit = 25000000
	case "platinum":
		tier = l.Tier
		limit = 50000000
	default:
		tier = "bronze"
		limit = 5000000
	}

	return tier, limit
}
