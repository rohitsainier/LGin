package main

import (
	"LGin/controller"
	"LGin/middlewares"
	"LGin/service"
	"io"
	"net/http"
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
	//setting up log output
	SetupLogOutput()

	server := gin.New()

	//loading ui
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	//middlewares
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", videos)
		apiRoutes.POST("/videos/add", saveVideo)
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}

func videos(ctx *gin.Context) {
	ctx.JSON(200, videoController.FindAll())
}
func saveVideo(ctx *gin.Context) {
	err := videoController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Video added successfully!",
		})
	}

}
