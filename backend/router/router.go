package router

import (
	"backend/controllers/category"
	"backend/controllers/courses"
	"backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {

	engine.GET("/users/:id", users.GetUserById)
	engine.POST("/users", users.CreateUser)
	engine.POST("/users/login", users.Login)
	engine.GET("/users/token/:token", users.GetUsersByToken)
	//engine.PUT("/users/:id", users.UpdateUser)

	engine.GET("/courses/:id", courses.GetCourseById)
	engine.POST("/courses", courses.CreateCourse)
	engine.PUT("/courses/update/:id", courses.UpdateCourse)
	engine.PUT("/courses/delete/:id", courses.DeleteCourse)
	//engine.PUT("/courses/:id", courses.UpdateCourse)

	engine.GET("/users/courses/:id", users.GetUserCourses)

	engine.GET("/category/:id", category.GetCategoryById)
	engine.POST("/category", category.CreateCategory)

}
