package main

import (
	_ "CRUD-GO/docs"
	"CRUD-GO/src/configuration/database/mongodb"
	"CRUD-GO/src/configuration/logger"
	"CRUD-GO/src/controller/routes"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title CRUD-GO | Rafael Freire
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {
	logger.Info("About to start user application")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
