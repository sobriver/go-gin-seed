package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-seed/pkg/app"
	"go-gin-seed/pkg/e"
	"go-gin-seed/pkg/logger"
	"go-gin-seed/routers/request"
	"go-gin-seed/service/user_service"
)

// user login
func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.LoginForm
	)
	logger.Info("user login form:%v", form)
	//check param
	errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil)
		return
	}
	//exec service
	errCode = user_service.Login(form)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil)
		return
	}
	//return success response
	appG.Response(e.SUCCESS, nil)

}
