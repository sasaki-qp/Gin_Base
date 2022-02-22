package user_controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"tutorial/model"
	"tutorial/service/user_service"
)

func GetUser(ctx *gin.Context) {
	// var userService user_service.UserService = user_service.UserService{}
	var userService user_service.UserService
	uid, err := strconv.Atoi(ctx.Param("uid"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": 500,
			"message": "Internal server error",
		})
		return
	} 
	selectedUser, err := userService.GetUser(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": 500,
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"user": selectedUser,
	})
	return;
}

func PostUser(ctx * gin.Context) {
	var requestBody user_model.PostRequestBody
	if err := ctx.ShouldBindJSON(&requestBody); 
		err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"statusCode": 400,
				"message": "Bad Request",
			})
			return;
		}
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"isTest": requestBody.IsTest,
		"uid": requestBody.Uid,
		"uname": requestBody.Uname,
	})
	return;
}