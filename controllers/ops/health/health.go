// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package health

import (
	"app/constants"
	"app/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/e"
	"github.com/ysicing/ext/logger"
)

// @Summary health
// @version 0.0.1
// @Accept application/json
// @Tags Health
// @Success 200
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(200, e.Done("I am ok."))
}

// @Summary version
// @version 0.0.1
// @Accept application/json
// @Tags Health
// @Success 200
// @Router /version [get]
func RVersion(c *gin.Context) {
	c.JSON(200, e.Done(map[string]string{
		"builddate": constants.Date,
		"release":   constants.Release,
		"gitcommit": constants.Commit,
	}))
}

// @Summary errpage
// @version 0.0.1
// @Accept application/json
// @Tags Health
// @Success 500
// @Router /err500 [get]
func Err500(c *gin.Context) {
	logger.Slog.Error("too long err")
	c.JSON(500, e.Error(10500, "500 Err by Gins!"))
}

// @Summary errpanic
// @version 0.0.1
// @Accept application/json
// @Tags Health
// @Success 500
// @Router /errpanic [get]
func ErrPanic(c *gin.Context) {
	panic("panic_err")
	c.JSON(500, e.Error(10500, "Test panic err by Gins!"))
}

// @Summary 生成测试Token
// @version 0.0.1
// @Tags Health
// @Accept application/json
// @Success 200
// @Router /gentoken [get]
func GenToken(c *gin.Context) {
	token, _ := jwt.JwtAuth("admin", "admin")
	c.JSON(200, e.Done(map[string]interface{}{
		"user":  "admin",
		"token": fmt.Sprintf("Bearer %v", token),
	}))
}
