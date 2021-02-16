// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exampleapi

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/exgin"
	"github.com/ysicing/ext/file"
	"github.com/ysicing/ext/gerr"
	"github.com/ysicing/ext/ztime"
)

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
		gerr.Dangerous("不支持mysql")
		return
	}
	dbres := file.FileSize2Str(dbdsn)
	if len(dbres) != 0 {
		exgin.GinsData(c, map[string]interface{}{
			"timestamp": ztime.NowFormat(),
			"size":      dbres,
		}, nil)
		return
	}
	gerr.Dangerous("文件不存在")
}
