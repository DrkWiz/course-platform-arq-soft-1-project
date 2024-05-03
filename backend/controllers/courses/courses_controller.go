package courses

import (
	coursesDomain "backend/domain/courses"
	coursesService "backend/services/courses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	var createCourseRequest coursesDomain.CreateCourseRequest
	c.BindJSON(&createCourseRequest)
	response := coursesService.CreateCourse(createCourseRequest)
	c.JSON(http.StatusOK, response)
}
