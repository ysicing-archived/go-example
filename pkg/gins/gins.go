// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package gins

import (
	"sync"

	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/exgin"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Gins *gin.Engine
var engineOnce sync.Once

func GinInit() {
	engineOnce.Do(func() {
		debug := viper.GetBool("server.debug")
		Gins = exgin.Init(&exgin.Config{
			Debug:   debug,
			Gops:    debug,
			Pprof:   debug,
			Cors:    true,
			Metrics: true,
		})
		Gins.Use(exgin.ExLog("/swagger"), exgin.ExRecovery(), exgin.ExTraceID())
		logrus.Debug(color.SGreen("create gin engine success..."))
	})
}
