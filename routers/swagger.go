// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package routers

import (
	_ "app/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	register("swagger", func(r *gin.Engine) {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	})
}
