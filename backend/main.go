package main

import (
    "backend/db"
    "backend/router"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors" // Importa el paquete cors
)

func main() {
    db.StartDbEngine()
    engine := gin.New()

    // Configura CORS
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true // Permite todas las solicitudes de origen
    engine.Use(cors.New(config)) // Usa el middleware CORS

    router.MapUrls(engine)
    engine.Run(":8080")
}