package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-gin-seed/pkg/e"
	"go-gin-seed/pkg/logger"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) int {
	err := c.Bind(form)
	if err != nil {
		logger.Error("bind error:%s", err)
		return e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		logger.Error("valid error:%s", err)
		return e.FAILURE
	}
	if !check {
		MarkErrors(valid.Errors)
		return e.INVALID_PARAMS
	}

	return e.SUCCESS
}

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logrus.Error(err.Key, err.Message)
	}
	return
}
