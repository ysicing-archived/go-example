// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package routers

import (
	"app/controllers/exampleapi"
	"github.com/gin-gonic/gin"
)

func init() {
	register("exampleapi", func(r *gin.Engine) {
		rg := r.Group("/api.example.com")
		rg.GET("/v1beta/db/total", exampleapi.DBTotal)
	})
}
