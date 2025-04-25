package controller

import (
	"CRUD-GO/src/configuration/rest_err"
	"CRUD-GO/src/controller/model/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

}
