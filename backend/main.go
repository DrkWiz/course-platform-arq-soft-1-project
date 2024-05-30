package main

import (
	"backend/db"
	"backend/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.StartDbEngine()
	engine := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	engine.Use(cors.New(config))
	router.MapUrls(engine)
	engine.Run(":8080")
}
