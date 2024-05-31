package users

import (
	"backend/dto"
	usersService "backend/services/users"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	e "backend/utils/errors"

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
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, token)
}

func GetUsersByToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	response, err := usersService.GetUsersByToken(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// UserCourses

func GetUserCourses(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	log.Print("GetUserCourses: ", id)

	response, err1 := usersService.GetUserCourses(id)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Inscribir a un usuario a un curso

func AddUserCourse(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if strings.Split(token, " ")[0] != "Bearer" {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}

	token = strings.Split(token, " ")[1]
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	err1 := usersService.AddUserCourse(id, token)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusCreated, "User added to course")
}

func GetUserCoursesByToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	response, err := usersService.GetUserCoursesByToken(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetIsAdmin(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Authorization header is required"))
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	response, err := usersService.GetIsAdmin(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Remove user from usercourse

func UnsubscribeUserCourse(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if strings.Split(token, " ")[0] != "Bearer" {
		c.JSON(http.StatusForbidden, "Forbidden")
		return
	}

	token = strings.Split(token, " ")[1]
	courseId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	err1 := usersService.UnsubscribeUserCourse(courseId, token)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusCreated, "User removed from course")
}
