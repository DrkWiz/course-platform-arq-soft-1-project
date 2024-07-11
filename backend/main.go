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

	// Update CORS configuration
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Include Authorization here
		AllowCredentials: true,
	}

	engine.Use(cors.New(config))
	router.MapUrls(engine)
	engine.Run(":8080")
}
