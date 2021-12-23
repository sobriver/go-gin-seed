package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"go-gin-seed/pkg/e"
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
		logrus.Error("select auth fail:{}", err)
		return nil, e.FAILURE
	}
	return &user, e.SUCCESS
}
