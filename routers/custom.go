// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package routers

import (
	"app/assets"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/e"
	"github.com/ysicing/ext/logger"
)

func init() {
	register("custom", func(r *gin.Engine) {
		r.GET("/too_long_err", func(c *gin.Context) {
			logger.Slog.Error("too long err")
			c.JSON(500, e.Error("10500", "Test 500 err by Gins!"))
		})
		r.GET("/panic_err", func(c *gin.Context) {
			panic("panic_err")
			c.JSON(500, e.Error("10500", "Test panic err by Gins!"))
		})
		// noroute
		r.NoRoute(func(c *gin.Context) {
			c.FileFromFS("/", assets.EmbedFS())
		})
		r.NoMethod(func(c *gin.Context) {
			c.JSON(404, e.Error("10404", c.Request.Method))
		})
	})
}
