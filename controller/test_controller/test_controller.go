package test_controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"net/http"
	"math/rand"
)

// channelを使用して何秒寝てたかメインゴールーチンに知らせたい
func Sleeping(c chan<- int) {
	log.Print("sleeping ...")
	rand.Seed(time.Now().Unix())
	var num int = rand.Intn(5)
	// func Sleep(d Duration): 現在のゴールーチンを少なくとも期間dの間停止させる
	time.Sleep(time.Duration(num) * time.Second)
	log.Print("just woke up !")
	c <- num //チャネルcに値を送信
}

func Concurrency(ctx *gin.Context) {
	c := make(chan int)   // チャネル作成
	go Sleeping(c)        // チャネルcを引数にゴールーチン起動
	var seconds int = <-c // 値が格納されているチェネルcから値を取り出す
	close(c)              //チャネルクローズ
	log.Print(seconds, "ミリ秒寝てました")
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"message":    "success",
	})
	return
}
