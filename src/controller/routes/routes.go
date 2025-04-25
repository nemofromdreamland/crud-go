package routes

import (
	"CRUD-GO/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

	r.POST("/login")

	r.GET("/swagger/*any")
}
