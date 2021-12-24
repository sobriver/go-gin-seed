package user_service

import (
	"go-gin-seed/models"
	"go-gin-seed/pkg/e"
	"go-gin-seed/routers/request"
)

// user login
func Login(form request.LoginForm) int {
	user := models.User{
		UserName: form.LoginName,
		Password: form.Password,
	}

	_, errCode := user.SelectByAuth()
	if errCode != e.SUCCESS {
		return errCode
	}
	return e.SUCCESS
}
