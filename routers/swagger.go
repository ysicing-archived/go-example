// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "app/docs"
)

func init() {
	register("swagger", func(r *gin.Engine) {
		url := ginSwagger.URL("http://127.0.0.1:7070/swagger/doc.json")
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	})
}