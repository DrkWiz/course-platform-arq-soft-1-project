package courses

import (
	"backend/dto"
	s "backend/services/courses"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	e "backend/utils/errors"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Get Course by ID

func GetCourseById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	log.Print("GetCourseById: ", id)

	response, err1 := s.CoursesService.GetCourseById(id)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response)
}

//Create new course

func CreateCourse(c *gin.Context) {
	var course dto.CourseCreateDto
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.CoursesService.CreateCourse(course)
	c.JSON(http.StatusOK, "Course created")
}

// Update course

func UpdateCourse(c *gin.Context) {
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

	var course dto.CourseUpdateDto
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, e.NewBadRequestApiError("Invalid JSON body"))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad ID")
		return
	}

	err1 := s.CoursesService.UpdateCourse(id, course, token)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusNoContent, "Course updated")

}

//Soft delete course

func DeleteCourse(c *gin.Context) {
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

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	err1 := s.CoursesService.DeleteCourse(id)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, "Course deleted")
}

// Get all courses in db

func GetCourses(c *gin.Context) {
	response, err1 := s.CoursesService.GetCourses()

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Check if the token is the owner of the course

func CheckOwner(c *gin.Context) {
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

	response, err1 := s.CoursesService.CheckOwner(token, courseId)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response)
}

func ImageUpload(c *gin.Context) {
	file, _ := c.FormFile("image")
	path := filepath.Join("./uploads", file.Filename)

	// Upload the file to a specific destination
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the relative path to be used in the frontend
	c.JSON(http.StatusOK, gin.H{"picture_path": file.Filename})
}

func GetImage(c *gin.Context) {
	picturepath := c.Param("picturepath")
	path := filepath.Join("./uploads", picturepath)

	c.File(path)
}
