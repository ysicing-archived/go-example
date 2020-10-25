// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/ginmid"
)

func init() {
	registerWithWeight("requestid", 80, func() gin.HandlerFunc {
		return ginmid.RequestID()
	})
}
