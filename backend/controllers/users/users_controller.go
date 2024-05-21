package users

import (
	usersDomain "backend/domain/users"
	usersService "backend/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//validate with db

	var loginRequest usersDomain.LoginRequest    // aca se crea la variable loginRequest de tipo LoginRequest
	c.BindJSON(&loginRequest)                    // aca se bindea el json que viene en el request a la variable loginRequest
	response := usersService.Login(loginRequest) // aca se llama a la funcion Login del paquete usersService y se le pasa como parametro la variable loginRequest
	c.JSON(http.StatusOK, response)              // aca se responde con un json con el status 200 y la variable response
}
