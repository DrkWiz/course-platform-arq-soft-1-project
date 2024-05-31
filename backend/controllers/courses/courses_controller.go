package courses

import (
	"backend/dto"
	coursesService "backend/services/courses"
	"net/http"
	"strconv"

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

	response, err1 := coursesService.CoursesService.GetCourseById(id)

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

	coursesService.CreateCourse(course)
	c.JSON(http.StatusOK, "Course created")
}

// Update course

func UpdateCourse(c *gin.Context) {
	var course dto.CourseUpdateDto
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad ID")
		return
	}

	err1 := coursesService.UpdateCourse(id, course)

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, "Course updated")

}

//Soft delete course

func DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	err1 := coursesService.DeleteCourse(id)

	if err1 != nil {
		c.JSON(http.StatusTeapot, err1)
		return
	}

	c.JSON(http.StatusOK, "Course deleted")
}

// Get all courses in db

func GetCourses(c *gin.Context) {
	response, err1 := coursesService.GetCourses()

	if err1 != nil {
		c.JSON(err1.Status(), err1)
		return
	}

	c.JSON(http.StatusOK, response)
}
