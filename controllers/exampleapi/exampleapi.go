// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package exampleapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/e"
	"github.com/ysicing/ext/utils/extime"
	"os"
)

// @Summary 查看DB大小
// @version 0.0.1
// @Accept application/json
// @Success 200
// @Router /api.example.com/v1beta/db/total [get]
func DBTotal(c *gin.Context) {
	dbtype := viper.GetString("db.type")
	dbdsn := viper.GetString("db.dsn")
	if dbtype == "mysql" {
		c.JSON(200, e.Error("不支持mysql"))
		return
	}
	fi, err := os.Stat(dbdsn)
	if err == nil {
		bs := float64(fi.Size())
		kbs := bs / 1024.0
		mbs := kbs / 1024.0
		c.JSON(200, e.Done(map[string]interface{}{
			"timestamp": extime.NowFormat(),
			"size":      fmt.Sprintf("%vM", mbs),
		}))
		return
	}
	c.JSON(200, e.Error("文件不存在"))
}