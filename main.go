package main

import (
	"CRUD-GO/src/configuration/database/mongodb"
	"CRUD-GO/src/configuration/logger"
	"CRUD-GO/src/controller"
	"CRUD-GO/src/controller/routes"
	"CRUD-GO/src/model/service"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()

	service := service.NewUserDomainService(nil)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
