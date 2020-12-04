// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/logger"
)

func ShowDebugMsg(s ...interface{}) {
	if viper.GetBool("server.debug") {
		logger.Slog.Debug(s...)
	}
}

// CheckAndExit check & exit
func CheckAndExit(err error) {
	if err != nil {
		logger.Slog.Fatal(err)
	}
}

// BindAndValid 校验参数
func BindAndValid(c *gin.Context, form interface{}) bool {
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Slog.Debug(form)
		logger.Slog.Errorf("err.bind: %v", err)
		return false
	}
	return true
}
