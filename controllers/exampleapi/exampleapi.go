// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package exampleapi

import (
	"github.com/ergoapi/errors"
	"github.com/ergoapi/exgin"
	"github.com/ergoapi/util/file"
	"github.com/ergoapi/util/ztime"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// DBTotal
// @Summary 查看DB大小
// @version 0.0.1
// @Tags 示例API
// @Accept application/json
// @Param Authorization header string true "token"
// @Security ApiKeyAuth
// @Success 200
// @Router /apis/example.dev/v1beta/db/total [get]
func DBTotal(c *gin.Context) {
	dbtype := viper.GetString("db.type")
	dbdsn := viper.GetString("db.dsn")
	if dbtype == "mysql" {
		errors.Dangerous("不支持mysql")
		return
	}
	dbres := file.Size2Str(dbdsn)
	if len(dbres) != 0 {
		exgin.GinsData(c, map[string]interface{}{
			"timestamp": ztime.NowFormat(),
			"size":      dbres,
		}, nil)
		return
	}
	errors.Dangerous("文件不存在")
}
