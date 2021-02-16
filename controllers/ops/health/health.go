// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package health

import (
	"app/constants"
	"app/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/exgin"
	"github.com/ysicing/ext/gerr"
	"github.com/ysicing/ext/logger/zlog"
)

// @Summary health
// @version 0.0.1
// @Accept application/json
// @Tags 默认
// @Success 200
// @Router /health [get]
func Health(c *gin.Context) {
	exgin.GinsData(c, "I am ok.", nil)
}

// @Summary version
// @version 0.0.1
// @Accept application/json
// @Tags 默认
// @Success 200
// @Router /version [get]
func RVersion(c *gin.Context) {
	exgin.GinsData(c, map[string]string{
		"builddate": constants.Date,
		"release":   constants.Release,
		"gitcommit": constants.Commit,
	}, nil)
}

// @Summary errpage
// @version 0.0.1
// @Accept application/json
// @Tags 默认
// @Success 500
// @Router /err500 [get]
func Err500(c *gin.Context) {
	zlog.Error("too long err")
	gerr.Bomb("500 Err by Gins!")
}

// @Summary errpanic
// @version 0.0.1
// @Accept application/json
// @Tags 默认
// @Success 500
// @Router /errpanic [get]
func ErrPanic(c *gin.Context) {
	panic("panic_err")
	gerr.Bomb("Test panic err by Gins!")
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
	exgin.Bind(c, &user)
	token, _ := jwt.JwtAuth(user.UserName, user.UserRole)
	data := map[string]interface{}{
		"user":  user,
		"token": fmt.Sprintf("Bearer %v", token),
	}
	exgin.GinsData(c, data, nil)
}
