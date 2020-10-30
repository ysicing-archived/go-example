// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package routers

import (
	"app/assets"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path"
	"strings"
)

var defaultrelativePaths = []string{"/css", "/js", "/fonts"}

func init() {
	register("ui", func(r *gin.Engine) {
		rg := r.Group("/")
		fs := assets.EmbedFS()
		relativePaths := viper.GetStringSlice("ui.static_path")
		if len(relativePaths) == 0 {
			relativePaths = defaultrelativePaths
		}
		handler := func(c *gin.Context) {
			c.FileFromFS(strings.TrimPrefix(c.Request.URL.Path, rg.BasePath()), fs)
		}
		for _, relativePath := range relativePaths {
			urlPattern := relativePath
			if urlPattern != "/" {
				urlPattern = path.Join(relativePath, "/*filepath")
			}
			rg.GET(urlPattern, handler)
			rg.HEAD(urlPattern, handler)
		}

	})
}
