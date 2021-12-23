package user_service

import (
	"go-gin-seed/models"
	"go-gin-seed/pkg/e"
	v1 "go-gin-seed/routers/api/v1"
)

// Login
func Login(form v1.LoginForm) int {
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
