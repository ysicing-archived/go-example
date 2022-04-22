// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package gins

import (
	"fmt"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/ergoapi/errors"
	"github.com/ergoapi/exgin"
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/exid"
	"github.com/ergoapi/util/file"
	"github.com/ergoapi/util/ztime"
	"github.com/ergoapi/zlog"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const headerXRequestID = "X-Request-ID"

var Gins *gin.Engine
var engineOnce sync.Once

func GinInit() {
	engineOnce.Do(func() {
		debug := viper.GetBool("server.debug")
		Gins = exgin.Init(debug)
		Gins.Use(exgin.ExCors(), ExRid(), ExLog(), Exrecovery())
		zlog.Debug(color.SGreen("create gin engine success..."))
	})
}

// ExRid rid 请求ID
func ExRid() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(headerXRequestID)
		if requestID == "" {
			requestID = exid.GenUUID()
		}
		c.Set(headerXRequestID, requestID)
		c.Writer.Header().Set(headerXRequestID, requestID)
		c.Next()
	}
}

// ExLog exlog middleware
func ExLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		if len(query) == 0 {
			query = " - "
		}
		if latency > time.Second*1 {
			zlog.Warn("[msg] api %v query %v", path, latency)
		}
		if len(c.Errors) > 0 || c.Writer.Status() >= 500 {
			msg := fmt.Sprintf("requestid %v => %v | %v | %v | %v | %v | %v <= err: %v", exgin.GetRID(c), color.SRed("%v", c.Writer.Status()), c.ClientIP(), c.Request.Method, path, query, latency, c.Errors.String())
			zlog.Warn(msg)
			go file.Writefile(fmt.Sprintf("/tmp/%v.errreq.txt", ztime.GetToday()), msg)
		} else {
			zlog.Info("requestid %v => %v | %v | %v | %v | %v | %v ", exgin.GetRID(c), color.SGreen("%v", c.Writer.Status()), c.ClientIP(), c.Request.Method, path, query, latency)
		}
	}
}

// Exrecovery recovery
func Exrecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if res, ok := err.(errors.ErgoError); ok {
					exgin.GinsData(c, nil, fmt.Errorf(res.Message))
					c.Abort()
					return
				}
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zlog.Error("Recovery from brokenPipe ---> path: %v, err: %v, request: %v", c.Request.URL.Path, err, string(httpRequest))
					c.AbortWithStatusJSON(200, gin.H{
						"data":      nil,
						"message":   "请求broken",
						"timestamp": ztime.NowUnix(),
						"code":      10500,
					})
				} else {
					zlog.Error("Recovery from panic ---> err: %v, request: %v, stack: %v", err, string(httpRequest), string(debug.Stack()))
					c.AbortWithStatusJSON(200, gin.H{
						"data":      nil,
						"message":   "请求panic",
						"timestamp": ztime.NowUnix(),
						"code":      10500,
					})
				}
				return
			}
		}()
		c.Next()
	}
}
