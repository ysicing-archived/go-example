// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package gins

import (
	"app/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/utils/extime"
)

// BindAndValid 校验参数
func BindAndValid(c *gin.Context, form interface{}) {
	errors.Dangerous(c.ShouldBindJSON(form))
}

// done done
func respdone(data interface{}) gin.H {
	return gin.H{
		"data":      data,
		"timestamp": extime.NowUnix(),
		"code":      200,
	}
}

// error error
func resperror(code int64, data interface{}) gin.H {
	return gin.H{
		"message":   data,
		"timestamp": extime.NowUnix(),
		"code":      code,
	}
}

func renderMessage(c *gin.Context, code int64, v interface{}) {
	if v == nil {
		c.JSON(200, respdone(nil))
		return
	}

	switch t := v.(type) {
	case string:
		c.JSON(200, resperror(code, t))
	case error:
		c.JSON(200, resperror(code, t.Error()))
	}
}

func GinsData(c *gin.Context, data interface{}, code int64, err error) {
	if err == nil {
		c.JSON(200, respdone(data))
		return
	}

	renderMessage(c, code, err.Error())
}
