package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yash0412/gintest/entity"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}
