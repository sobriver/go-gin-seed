package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-seed/models"
	"go-gin-seed/pkg/logger"
	"go-gin-seed/pkg/setting"
	"go-gin-seed/routers"
	"net/http"
)

func init() {
	setting.Init()
	logger.Init()
	models.Init()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	ginHandler := routers.InitGin()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        ginHandler,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	logger.Info("run server port:%s", endPoint)
	server.ListenAndServe()
}
