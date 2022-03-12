package types

import "time"

type CurrencyCode string
type AccountCategory string

const (
	USD CurrencyCode = "USD"
	EUR CurrencyCode = "EUR"
)

const (
	Banks          AccountCategory = "banks"
	PaymentMethods AccountCategory = "payment_methods"
)

type Transaction struct {
	ID          string          `json:"id"`
	Description string          `json:"source"`
	Category    AccountCategory `json:"category"`
	Amount      float32         `json:"amount"`
	MadeOn      time.Time       `json:"date"`
}

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
	Total        int64         `json:"total"`
}
