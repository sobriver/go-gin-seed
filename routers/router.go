package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-seed/pkg/logger"
)

func InitGin() *gin.Engine {
	r := gin.New()
	r.Use(logger.LogHandler())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
	}

	return r
}
