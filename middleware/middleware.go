package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
    "time"
	"net/http"
    "go.uber.org/zap"
)

func CustomMiddleware(ctx *gin.Context){
	//request 処理の前
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}
	requestTime := time.Now()
	authorization := ctx.GetHeader("authorization")
	if authorization != "test" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"statusCode": 401,
			"message": "Unauthorized",
		})
	}
	ctx.Next()
	//request 処理の後
	logger.Info("Incoming request",
		// 外部パッケージからは小文字で始まるフィールド名にアクセスできない
		// zap.String("path", ctx.fullPath()),
		zap.Int("statusCode", ctx.Writer.Status()),
		zap.String("requestPath", ctx.FullPath()),
		zap.String("authorization", authorization),
		zap.Duration("elapsed", time.Now().Sub(requestTime)),
    )
}