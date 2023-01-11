// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package health

import (
	"app/constants"
	"github.com/ergoapi/util/exjwt"
	"fmt"

	"github.com/ergoapi/util/exgin"
	"github.com/gin-gonic/gin"
)

// Health
//
//	@Summary health
//
// @version 0.0.1
// @Accept pplication/json
// @Tags 默认
// @Success 200
// @Router /health [get]
func Health(c *gin.Context) {
	exgin.GinsData(c, "I am ok.", nil)
}

// RVersion
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

type User struct {
	UserName string `json:"username" form:"username"`
	UserRole string `json:"userrole" form:"userrole"`
}

// GenToken
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
	token, _ := exjwt.Auth(user.UserName, user.UserRole)
	data := map[string]interface{}{
		"user":  user,
		"token": fmt.Sprintf("Bearer %v", token),
	}
	exgin.GinsData(c, data, nil)
}
