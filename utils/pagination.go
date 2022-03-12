package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get and return limit, offset based on request
func HandlePagination(c *gin.Context) (int, int) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		fmt.Errorf("error limit")
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		fmt.Errorf("error page")
	}
	offset := (page - 1) * limit

	return limit, offset
}
