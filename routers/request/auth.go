package request

import "go-gin-seed/pkg/app"

//user login form
type LoginRequest struct {
	app.BaseRequest
	LoginName string `form:"loginName" binding:"required,min=2,max=100"`
	Password  string `form:"password" binding:"required,min=2,max=100"`
}
