package models

import (
	"github.com/jinzhu/gorm"
	"go-gin-seed/pkg/e"
	"go-gin-seed/pkg/logger"
)

type User struct {
	Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
	State    int    `json:"state"`
}

func (u *User) SelectByAuth() (*User, int) {
	var user User
	err := db.Where("user_name = ? AND password = ? ", u.UserName, u.Password).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, e.PWD_NOT_MATCH
	}
	if err != nil {
		logger.Error("select auth fail:%s", err)
		return nil, e.FAILURE
	}
	return &user, e.SUCCESS
}
