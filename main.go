// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"app/cmd"
	"github.com/ysicing/ext/logger"
)

func init() {
	logcfg := &logger.LogConfig{Simple: true}
	logger.InitLogger(logcfg)
}

func main() {
	cmd.Execute()
}
