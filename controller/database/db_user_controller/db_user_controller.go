package db_user_controller

import(
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"net/http"
	"tutorial/model"
	"tutorial/service/database/user_service"
)

func FindUser(ctx *gin.Context) {
	var UserService user_service.DBUserService
	var user user_model.Users

	id, strerr := strconv.Atoi(ctx.Param("id"))
	if strerr != nil {
		log.Print("DEBUG strconv error", strerr)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": 500,
			"message": "Internal server error",
		})
		return
	}

	var err error = UserService.FindUser(&user, id)
	if err != nil {
		log.Print("DEBUG Database find error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": 500,
			"message": "Internal server error",
		})
		return
	}
	log.Print("DEBUG: successfully find user", user)
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"user": user,
	})
	return
	
} 