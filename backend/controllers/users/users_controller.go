package users

import (
	"backend/dto"
	usersService "backend/services/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

// Get Student by ID

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	log.Print("GetUserById: ", id)

	response, err1 := usersService.UsersService.GetUserById(id)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response)
}

//Create new user

func CreateUser(c *gin.Context) {
	var user dto.UserCreateDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err1 := usersService.CreateUser(user)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, "User created")
}

// Login

func Login(c *gin.Context) {
	var user dto.UserLoginDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err1 := usersService.Login(user)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, token)
}

func GetUsersByToken(c *gin.Context) {
	token := c.Param("token")

	response, err1 := usersService.GetUsersByToken(token)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response)
}
