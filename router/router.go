package router

import (
	"github.com/gin-gonic/gin"
	"tutorial/controller/user_controller"
	"tutorial/controller/test_controller"
	"tutorial/controller/database/db_user_controller"
)

func Routing(router *gin.Engine) {
	baseGroup := router.Group("/apis/v1")
	{
		userGroup := baseGroup.Group("/user")
		{
			userGroup.GET("/:uid", user_controller.GetUser)
			userGroup.POST("/", user_controller.PostUser)
		}
		testGroup := baseGroup.Group("/test")
		{
			testGroup.GET("/", test_controller.Concurrency)
		}
		dbUserGroup := baseGroup.Group("/dbuser")
		{
			dbUserGroup.GET("/:id", db_user_controller.FindUser)
		}
	}
}