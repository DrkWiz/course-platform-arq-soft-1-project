package courses

import (
	coursesDomain "backend/domain/courses"
	coursesService "backend/services/courses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	var createCourseRequest coursesDomain.CourseCreateRequest    // aca se crea la variable createCourseRequest de tipo CourseCreateRequest
	c.BindJSON(&createCourseRequest)                             // aca se bindea el json que viene en el request a la variable createCourseRequest
	response := coursesService.CreateCourse(createCourseRequest) // aca se llama a la funcion CreateCourse del paquete coursesService y se le pasa como parametro la variable createCourseRequest
	c.JSON(http.StatusCreated, response)                         // aca se responde con un json con el status 201 y la variable response
}
