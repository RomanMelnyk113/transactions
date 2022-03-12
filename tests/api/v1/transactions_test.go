package tests

import (
	"encoding/json"
	"fmt"
	"juni/task/configs"
	"juni/task/routers"
	"juni/task/types"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func init() {
	router = routers.InitRoutes()
	configs.LoadEnv()
	// TODO: use test db instead of app db
	configs.SetupDB()
}

func TestTransactions(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/transactions", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response types.TransactionsResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		panic(err)
	}
	var total int64 = 1
	assert.Equal(t, total, response.Total)
	assert.Equal(t, transaction_0.ID, response.Transactions[0].ID)
	assert.Equal(t, transaction_0.Description, response.Transactions[0].Description)
	assert.Equal(t, transaction_0.Category, response.Transactions[0].Category)
	assert.Equal(t, transaction_0.Amount, response.Transactions[0].Amount)
}

func TestSubTransactions(t *testing.T) {
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/api/v1/transactions/%s/sublist", transaction_0.ID)
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response types.TransactionsResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		panic(err)
	}
	var total int64 = 4
	assert.Equal(t, total, response.Total)
	assert.Equal(t, transaction_1.ID, response.Transactions[0].ID)
	assert.Equal(t, transaction_2.ID, response.Transactions[1].ID)
	assert.Equal(t, transaction_3.ID, response.Transactions[2].ID)
	assert.Equal(t, transaction_4.ID, response.Transactions[3].ID)
}

func TestSubTransactionsPagination(t *testing.T) {
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/api/v1/transactions/%s/sublist?limit=2&page=2", transaction_0.ID)
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response types.TransactionsResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		panic(err)
	}
	var total int64 = 4
	assert.Equal(t, total, response.Total)
	assert.Equal(t, transaction_3.ID, response.Transactions[0].ID)
	assert.Equal(t, transaction_4.ID, response.Transactions[1].ID)
}
