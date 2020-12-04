// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package health

import (
	"app/constants"
	"app/pkg/jwt"
	"app/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/e"
	"github.com/ysicing/ext/logger"
)

// @Summary health
// @version 0.0.1
// @Accept application/json
// @Tags 默认
// @Success 200
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(200, e.Done("I am ok."))
}

// @Summary version
// @version 0.0.1
// @Accept application/json
// @Tags 默认
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
// @Tags 默认
// @Success 500
// @Router /err500 [get]
func Err500(c *gin.Context) {
	logger.Slog.Error("too long err")
	c.JSON(500, e.Error(10500, "500 Err by Gins!"))
}

// @Summary errpanic
// @version 0.0.1
// @Accept application/json
// @Tags 默认
// @Success 500
// @Router /errpanic [get]
func ErrPanic(c *gin.Context) {
	panic("panic_err")
	c.JSON(500, e.Error(10500, "Test panic err by Gins!"))
}

type User struct {
	UserName string `json:"username" form:"username"`
	UserRole string `json:"userrole" form:"userrole"`
}

// @Summary 生成测试Token
// @version 0.0.1
// @Tags 默认
// @Accept application/json
// @Param user body User true "用户信息"
// @Success 200
// @Router /gentoken [post]
func GenToken(c *gin.Context) {
	var user User
	if !utils.BindAndValid(c, &user) {
		c.JSON(200, e.Error(10400, "校验数据失败"))
		return
	}
	token, _ := jwt.JwtAuth(user.UserName, user.UserRole)
	c.JSON(200, e.Done(map[string]interface{}{
		"user":  user,
		"token": fmt.Sprintf("Bearer %v", token),
	}))
}
