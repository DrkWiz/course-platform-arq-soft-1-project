package router

import (
	"backend/controllers/courses"
	"backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	engine.POST("/users/login", users.Login)

	engine.POST("/courses/create", courses.CreateCourse)
	engine.GET("/students/:id", users.GetStudentById)

}
