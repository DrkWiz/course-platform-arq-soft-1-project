package users

import (
	usersDomain "backend/domain/users"
	usersService "backend/services/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	//validate with db

	var loginRequest usersDomain.LoginRequest
	c.BindJSON(&loginRequest)
	response := usersService.Login(loginRequest)
	c.JSON(http.StatusOK, response)
}

// Get Student by ID

func GetStudentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	log.Print("GetStudentById: ", id)

	response, err1 := usersService.UsersService.GetStudentById(id)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response)
}
