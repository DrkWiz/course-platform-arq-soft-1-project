package router

import (
	"backend/controllers/courses"
	"backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {

	engine.GET("/users/:id", users.GetUserById)
	engine.POST("/users", users.CreateUser)

	engine.GET("/courses/:id", courses.GetCourseById)
	engine.POST("/courses", courses.CreateCourse)
	engine.PUT("/courses/update/:id", courses.UpdateCourse)

	engine.PUT("/courses/delete/:id", courses.DeleteCourse)
}
