package v1

import (
	"juni/task/repositories"
	"juni/task/utils"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// Load top level transactions
// @Summary Top level transactions
// @Schemes
// @Description Returns top level transactions
// @Tags transactions
// @Produce json
// @Param limit query int 10 "Items per page limit"
// @Param page query int 1 "Page number"
// @Success 200 {object} types.TransactionsResponse
// @Router /transactions [get]
func GetTransactions(c *gin.Context) {
	limit, offset := utils.HandlePagination(c)
	transactions, count := repositories.LoadTransactions(limit, offset)

	c.JSON(200, gin.H{
		"transactions": transactions,
		"total":        count,
	})
}

// @BasePath /api/v1

// Load related transactions based on transaction id
// @Summary Top level transactions
// @Schemes
// @Description Returns related transactions
// @Tags transactions_sub_list
// @Produce json
// @Param id path string nil "Transaction id"
// @Param limit query int 10 "Items per page limit"
// @Param page query int 1 "Page number"
// @Success 200 {object} types.TransactionsResponse
// @Router /transactions/{id}/sublist [get]
func GetTransactionsSublist(c *gin.Context) {
	transactionId := c.Param("id")
	limit, offset := utils.HandlePagination(c)

	transactions, count := repositories.LoadSubTransactions(transactionId, limit, offset)
	c.JSON(200, gin.H{
		"transactions": transactions,
		"total":        count,
	})
}
