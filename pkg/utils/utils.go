// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package utils

import (
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