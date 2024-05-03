package users

import (
	usersDomain "backend/domain/users"
	usersService "backend/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//validate with db

	var loginRequest usersDomain.LoginRequest
	c.BindJSON(&loginRequest)
	response := usersService.Login(loginRequest)
	c.JSON(http.StatusOK, response)
}
