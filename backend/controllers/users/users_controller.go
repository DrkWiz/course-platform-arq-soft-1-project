package users

import (
	usersDomain "backend/domain/users"
	usersService "backend/services/users"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//validate with db

	var loginRequest usersDomain.LoginRequest

	c.BindJSON(&loginRequest)
	usersService.Login(loginRequest)

}
