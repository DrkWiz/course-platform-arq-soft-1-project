package router

import (
	"backend/controllers/category"
	"backend/controllers/courses"
	"backend/controllers/users"
	"backend/services/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	// Public routes
	engine.POST("/users", users.CreateUser)
	engine.POST("/users/login", users.Login)
	engine.GET("/courses/:id", courses.GetCourseById)
	engine.GET("/category/:id", category.GetCategoryById)

	// Protected routes
	protected := engine.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/users/me", users.GetUsersByToken)
	protected.GET("/users/:id", users.GetUserById)
	protected.GET("/users/courses/:id", users.GetUserCourses)
	protected.GET("/users/courses/", users.GetUserCoursesByToken)
	protected.POST("/users/courses/:id", users.AddUserCourse)
	protected.POST("/courses", courses.CreateCourse)
	protected.PUT("/courses/update/:id", courses.UpdateCourse)
	protected.PUT("/courses/delete/:id", courses.DeleteCourse)
	protected.POST("/category", category.CreateCategory)
}
