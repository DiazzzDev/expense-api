package model

import (
	"github.com/shopspring/decimal"
)

type Expense struct {
	ID         int64           `json:"id"`
	Title      string          `json:"title"`
	Amount     decimal.Decimal `json:"amount"`
	CategoryId Category        `json:"category_id"`
}
