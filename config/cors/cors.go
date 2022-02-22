package cors

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func CustomCors() gin.HandlerFunc {
	var config cors.Config = cors.Config {
		// AllowOrigins: []string {
		// 	"https://example.com",
		// 	"https://example2.com",
		// },
		AllowAllOrigins: true,
		AllowCredentials: true,
		AllowHeaders: []string {
			"Origin",
			"Content-Type",
			"Content-Length",
		},
		AllowMethods: []string {
			"GET",
			"POST",
		},
	}

	return cors.New(config)
}