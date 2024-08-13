package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// /ping にアクセスがあった場合、"pong" という JSON を返す
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Runメソッドでサーバーを立てる
	r.Run("localhost:8080")
}
