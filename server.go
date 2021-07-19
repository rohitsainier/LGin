package main

import (
	"LGin/controller"
	"LGin/middlewares"
	"LGin/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func SetupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	SetupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/videos", videos)
	server.POST("/videos/add", saveVideo)
	server.Run(":8080")
}

func videos(ctx *gin.Context) {
	ctx.JSON(200, videoController.FindAll())
}
func saveVideo(ctx *gin.Context) {
	ctx.JSON(200, videoController.Save(ctx))
}
