package routers

import (
	v1 "juni/task/routers/api/v1"

	"github.com/gin-gonic/gin"

	docs "juni/task/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1Group := router.Group("/api/v1")
	v1Group.GET("/transactions", v1.GetTransactions)
	v1Group.GET("/transactions/:id/sublist", v1.GetTransactionsSublist)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
