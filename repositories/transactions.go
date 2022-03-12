package repositories

import (
	"juni/task/configs"
	"juni/task/types"

	"gorm.io/gorm"
)

// Returns base query to transactions table. Select count(*) or relevant fields depends on is_count param
func getBaseQuery(is_count bool) *gorm.DB {
	//SELECT description, made_on, amount, category
	//FROM transactions
	//LEFT JOIN accounts on accounts.id = transactions.account_id
	//WHERE is_parent = ?

	q := configs.DB.Debug().Table("transactions")
	if is_count {
		// NOTE: for better performance its better to use `count(*) over() as total`
		// but in current case let's use separate query to simplify code logic
		q = q.Select("count(*)")
	} else {
		q = q.Select("transactions.id, description, made_on, amount, accounts.category")
	}
	q = q.Joins("left join accounts on accounts.id = transactions.account_id")

	if !is_count {
		q = q.Order("transactions.order asc")
	}

	return q
}

// Returns transactions and total count of top level transactions
func LoadTransactions(limit int, offset int) ([]types.Transaction, int64) {
	var transactions []types.Transaction
	var count int64

	getBaseQuery(false).
		Where("parent_id is null").
		Limit(limit).
		Offset(offset).
		Scan(&transactions)

	getBaseQuery(true).
		Where("parent_id is null").
		Scan(&count)

	return transactions, count
}

// Returns transactions based on parent_id
func LoadSubTransactions(transactionId string, limit int, offset int) ([]types.Transaction, int64) {
	var transactions []types.Transaction
	var count int64

	getBaseQuery(false).
		Where("transactions.parent_id = ?", transactionId).
		Limit(limit).
		Offset(offset).
		Scan(&transactions)

	getBaseQuery(true).
		Where("transactions.parent_id = ?", transactionId).
		Scan(&count)

	return transactions, count
}
