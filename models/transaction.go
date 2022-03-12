package models

import (
	"juni/task/types"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Transaction data with details received from 3rd party service
type Transaction struct {
	gorm.Model
	ID           string             `json:"id"`
	AccountID    string             `json:"account_id"`
	MadeOn       time.Time          `json:"made_on"`
	Amount       float32            `json:"amount" sql:"type:decimal(10,2);"`
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Description  string             `json:"description"`
	Details      datatypes.JSON     `json:"details"`   // contains transaction details received as callback
	Order        uint               `json:"order"`     // transaction order to be displayed in the sublist
	ParentId     *string            `json:"parent_id"` // determines if it parent transaction or related to another trx

	// we can store category here to avoid additional joins, but in current case let's use it from Account
	// Category     AccountCategory `json:"category"`

	Account Account
	Parent  *Transaction
}

// User payment/bank account connected to the system
type Account struct {
	gorm.Model
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	CreatedAt time.Time             `json:"created_at"`
	Category  types.AccountCategory `json:"category"` // can be enum?
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
