package main

import (
	"LGin/controller"
	"LGin/middlewares"
	"LGin/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
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
