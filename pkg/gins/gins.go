// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package gins

import (
	"github.com/gin-gonic/gin"
	"github.com/google/gops/agent"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/exmisc"
	"sync"
)

var Gins *gin.Engine
var engineOnce sync.Once

func GinInit() {
	engineOnce.Do(func() {
		if viper.GetBool("server.debug") {
			gin.SetMode(gin.DebugMode)
			agentaddr := viper.GetString("server.agent")
			if len(agentaddr) == 0 {
				agentaddr = "0.0.0.0:8848"
			}
			go agent.Listen(agent.Options{
				Addr:            agentaddr,
				ShutdownCleanup: true})
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		gin.DisableConsoleColor()
		Gins = gin.New()
		logger.Slog.Info(exmisc.SGreen("create gin engine success..."))
	})
}
