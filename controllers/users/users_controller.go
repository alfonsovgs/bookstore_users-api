package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alfonsovgs/bookstore_users-api/domain/users"
	"github.com/alfonsovgs/bookstore_users-api/services"
	"github.com/alfonsovgs/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())

		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("User id should be a number")
		c.JSON(err.Status, err)
		return
	}

	result, errGet := services.GetUser(userId)

	if errGet != nil {
		c.JSON(errGet.Status, errGet)
		return
	}

	c.JSON(http.StatusOK, result)
}
