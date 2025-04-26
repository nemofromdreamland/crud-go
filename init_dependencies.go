package main

import (
	"CRUD-GO/src/controller"
	"CRUD-GO/src/model/repository"
	"CRUD-GO/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
