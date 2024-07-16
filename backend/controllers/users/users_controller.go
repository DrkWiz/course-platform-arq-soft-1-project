package users

import (
	"backend/dto"
	s "backend/services/users"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	e "backend/utils/errors"

	log "github.com/sirupsen/logrus"
)

// Get User by ID

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	log.Print("GetUserById: ", id)

	response, err1 := s.UsersService.GetUserById(id)

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

	err1 := s.UsersService.CreateUser(user)

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
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Invalid JSON body"))
		return
	}

	token, err1 := s.UsersService.Login(user)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, token)
}

// Obtiene el usuario del token

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

	response, err := s.UsersService.GetUsersByToken(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Obtiene los cursos a los que esta inscripto un usuario.

func GetUserCourses(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	log.Print("GetUserCourses: ", id)

	response, err1 := s.UsersService.GetUserCourses(id)

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
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	token = strings.Split(token, " ")[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	err1 := s.UsersService.AddUserCourse(id, token)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusCreated, "User added to course")
}

// Obtiene los cursos a los que esta inscripto un usuario con el token

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

	response, err := s.UsersService.GetUserCoursesByToken(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Revisa si el token es de un admin.

func CheckAdmin(c *gin.Context) {
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

	response, err := s.UsersService.CheckAdmin(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Elimina la inscripcion de un usuario a un curso.

func UnsubscribeUserCourse(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if strings.Split(token, " ")[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	token = strings.Split(token, " ")[1]
	if token == "" {
		c.JSON(http.StatusUnauthorized, e.NewUnauthorizedApiError("Token is required"))
		return
	}

	courseId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	err1 := s.UsersService.UnsubscribeUserCourse(courseId, token)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, "User removed from course")
}

// Devuelve si un usuario esta inscripto a un curso.

func CheckEnrolled(c *gin.Context) {
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

	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	response, err1 := s.UsersService.CheckEnrolled(courseId, token)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response)
}
