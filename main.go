package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// 有新的pr过来就发送通知到飞书

func main() {
	r := gin.Default()
	r.POST("/webhook", webhook)

	r.Run("localhost:9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func webhook(ctx *gin.Context) {
	fmt.Println(*ctx)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	
	fmt.Printf("ctx.Request.body: %v", string(data))
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("yes")
}
