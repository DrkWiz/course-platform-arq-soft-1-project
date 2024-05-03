package users

import (
	"github.com/gin-gonic/gin"

	"net/http"

	usersDomain "backend/domain/users"
	usersService "backend/services/users"
)

func Login(context *gin.Context) {
	var loginRequest usersDomain.LoginRequest
	context.BindJSON(&loginRequest)
	response := usersService.Login(loginRequest)
	context.JSON(http.StatusOK, response)
}
