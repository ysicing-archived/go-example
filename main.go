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

// @title Go Example API
// @version 0.0.1
// @description This is a sample server Petstore server.

// @contact.name ysicing
// @contact.url http://github.com/ysicing
// @contact.email i@ysicing.me

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:7070
func main() {
	cmd.Execute()
}
