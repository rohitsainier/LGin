package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/test", test)
	server.Run(":8080")
}

func test(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}
