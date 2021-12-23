package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-seed/pkg/app"
	"go-gin-seed/pkg/e"
)

type LoginForm struct {
	LoginName string `form:"loginName" valid:"Required;Max(100)"`
	Password  string `form:"password" valid:"Required;Max(100)"`
}

// Login /*
func Login(c *gin.Context) {
	//check param
	var (
		appG = app.Gin{C: c}
		form LoginForm
	)
	errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(errCode, e.GetMsg(errCode))
		return
	}
	//exec service

	//return success response
	appG.Response(e.SUCCESS, nil)

}
