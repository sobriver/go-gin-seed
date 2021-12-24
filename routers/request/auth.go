package request

import "go-gin-seed/pkg/app"

//user login form
type LoginForm struct {
	app.BaseRequest
	LoginName string `form:"loginName" valid:"Required;Max(100)"`
	Password  string `form:"password" valid:"Required;Max(100)"`
}
