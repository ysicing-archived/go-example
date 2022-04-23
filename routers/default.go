// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package routers

import (
	"app/controllers/ops/health"
	"fmt"

	"github.com/ergoapi/exgin"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	register("default.api", func(r *gin.Engine) {
		r.GET("/health", health.Health)
		r.GET("/version", health.RVersion)
		r.GET("/err500", health.Err500)
		r.GET("/errpanic", health.ErrPanic)
		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
		r.POST("/gentoken", health.GenToken)
		r.NoMethod(func(c *gin.Context) {
			exgin.GinsAbort(c, fmt.Sprintf("not fount: %v", c.Request.Method))
		})
		r.NoRoute(func(c *gin.Context) {
			exgin.GinsAbort(c, fmt.Sprintf("not fount: %v", c.Request.URL.Path))
		})
	})
}
