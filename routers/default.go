// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package routers

import (
	"app/controllers/ops/health"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/e"
)

func init() {
	register("ext.api", func(r *gin.Engine) {
		extapi := r.Group("/api.ext")

		v1 := extapi.Group("/v1")
		v1.GET("/health", health.Health)
		v1.GET("/version", health.RVersion)
		v1.GET("/err500", health.Err500)
		v1.GET("/errpanic", health.ErrPanic)

		r.NoMethod(func(c *gin.Context) {
			c.JSON(404, e.Error(10404, c.Request.Method))
		})
	})
}
