package router

import (
	"backend/controllers/category"
	"backend/controllers/courses"
	"backend/controllers/files"
	"backend/controllers/users"
	"backend/services/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	// Public routes
	engine.POST("/users", users.CreateUser)
	engine.POST("/users/login", users.Login)
	engine.GET("/courses/:id", courses.GetCourseById)
	engine.GET("/courses", courses.GetCourses)
	engine.GET("/category/:id", category.GetCategoryById)
	engine.GET("/category/all", category.GetCategories)

	engine.GET("/courses/:id/files", files.GetFilesByCourse)

	// Protected routes
	protected := engine.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/users/isAdmin", users.CheckAdmin)
	protected.GET("/users/me", users.GetUsersByToken)
	protected.GET("/users/:id", users.GetUserById)
	protected.GET("/users/courses/:id", users.GetUserCourses)
	protected.GET("/users/courses/all", users.GetUserCoursesByToken)
	protected.POST("/users/courses/:id", users.AddUserCourse)
	protected.DELETE("/users/courses/:id/unsubscribe", users.UnsubscribeUserCourse)
	protected.GET("/users/courses/:id/enrolled", users.CheckEnrolled)
	protected.GET("courses/:id/rating", courses.GetAvgRating)
	protected.GET("/courses/:id/comments", courses.GetComments)

	protected.POST("/courses", courses.CreateCourse)
	protected.POST("/courses/:id/owner", courses.CheckOwner)
	protected.PUT("/courses/update/:id", courses.UpdateCourse)
	protected.PUT("/courses/delete/:id", courses.DeleteCourse)
	protected.POST("/courses/:id/comments", courses.SetComment)

	protected.POST("/category", category.CreateCategory)

	protected.POST("/upload", courses.ImageUpload)
	protected.GET("/img/:picturepath", courses.GetImage)

	// Files
	protected.GET("/files/:id", files.GetFileById)
	protected.POST("/courses/:id/files", files.UploadFile)

	engine.Static("/uploads", "./uploads")

}
