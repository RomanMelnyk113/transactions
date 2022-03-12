package main

import (
	"juni/task/configs"
	"juni/task/models"
	"juni/task/routers"
)

//var DB *gorm.DB

func main() {
	configs.LoadEnv()
	configs.SetupDB()
	models.AutoMigration(configs.DB)

	router := routers.InitRoutes()
	router.Run()

}
