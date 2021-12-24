package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-seed/pkg/logger"
	v1 "go-gin-seed/routers/api/v1"
)

func InitGin() *gin.Engine {
	r := gin.New()
	r.Use(logger.LogHandler())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.POST("/auth/login", v1.Login)
	}

	return r
}
