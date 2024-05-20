package router

import (
	"backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {

	engine.GET("/users/:id", users.GetUserById)
	engine.POST("/users", users.CreateUser)
}
