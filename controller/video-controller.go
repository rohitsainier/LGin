package controller

import (
	"LGin/entity"
	"LGin/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.ResponseVideo
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.ResponseVideo {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	var response entity.ResponseVideo
	response.Code = 100
	response.Message = "Video added successfully"
	response.Data = video
	return response

}
