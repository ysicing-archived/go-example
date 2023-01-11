// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package routers

import (
	"app/controllers/exampleapi"

	"github.com/gin-gonic/gin"
)

func init() {
	register("exampleapi", func(r *gin.Engine) {
		rg := r.Group("/apis/example.dev")
		rg.GET("/v1beta/db/total", exampleapi.DBTotal)
		rg.POST("/v1beta/db/add", exampleapi.DBAdd)
	})
}
