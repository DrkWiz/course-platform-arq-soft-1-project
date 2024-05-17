package users

import (
	usersService "backend/services/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

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
