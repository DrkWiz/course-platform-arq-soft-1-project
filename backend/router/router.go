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
	//engine.GET("/courses/max/:id", courses.GetMaxCourseById)
	engine.GET("/courses", courses.GetCourses)
	engine.GET("/category/:id", category.GetCategoryById)

	// Protected routes
	protected := engine.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/users/isAdmin", users.CheckAdmin)
	protected.GET("/users/me", users.GetUsersByToken)
	protected.GET("/users/:id", users.GetUserById)
	protected.GET("/users/courses/:id", users.GetUserCourses)
	protected.GET("/users/courses/", users.GetUserCoursesByToken)
	protected.POST("/users/courses/:id", users.AddUserCourse)
	protected.DELETE("/users/courses/:id/unsubscribe", users.UnsubscribeUserCourse)
	protected.POST("/courses", courses.CreateCourse)
	protected.POST("/courses/:id/owner", courses.CheckOwner)
	protected.PUT("/courses/update/:id", courses.UpdateCourse)
	protected.PUT("/courses/delete/:id", courses.DeleteCourse)
	protected.POST("/category", category.CreateCategory)
	protected.POST("/upload", courses.ImageUpload)
	protected.GET("/img/:picturepath", courses.GetImage)
	engine.Static("/uploads", "./uploads")
}
