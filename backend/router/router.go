package router

import (
	"backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {

	engine.GET("/students/:id", users.GetStudentById)

}
