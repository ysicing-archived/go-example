// Copyright (c) 2023 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package exampleapi

import (
	"app/models"

	errors "github.com/ergoapi/util/exerror"
	"github.com/ergoapi/util/exgin"
	"github.com/ergoapi/util/exid"
	"github.com/ergoapi/util/expass"
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

// DBAdd
// @Summary 操作DB
// @version 0.0.1
// @Tags 示例API
// @Accept application/json
// @Param Authorization header string true "token"
// @Security ApiKeyAuth
// @Success 200
// @Router /apis/example.dev/v1beta/db/add [post]
func DBAdd(c *gin.Context) {
	user := exid.GenUUID()
	genuser := models.User{
		Username: user,
		Password: expass.PwGenAlphaNumSymbols(16),
		Email:    "",
		Banned:   false,
		Token:    exid.GenUID(user),
	}
	exgin.GinsData(c, genuser, genuser.Save())
}
