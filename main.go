package main

import (
	"github.com/gin-gonic/gin"
	"tutorial/router"
	"tutorial/middleware"
	"tutorial/config/cors"
	"tutorial/config/database"
)

func main() {
	engine := gin.Default()
	engine.Use(cors.CustomCors())
	engine.Use(middleware.CustomMiddleware)
	router.Routing(engine)
	database.InitDB()
	engine.Run(":8080")
}
