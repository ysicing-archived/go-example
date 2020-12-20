// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exampleapi

import (
	"app/pkg/errors"
	"app/pkg/gins"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/utils/exfile"
	"github.com/ysicing/ext/utils/extime"
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
		errors.Dangerous("不支持mysql")
		return
	}
	dbres := exfile.FileSize2Str(dbdsn)
	if len(dbres) != 0 {
		gins.GinsData(c, map[string]interface{}{
			"timestamp": extime.NowFormat(),
			"size":      dbres,
		}, 200, nil)
		return
	}
	errors.Dangerous("文件不存在")
}
