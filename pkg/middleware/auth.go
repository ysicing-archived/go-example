// Copyright (c) 2023 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package middleware

import (
	"app/pkg/rbac"
	"strings"

	"github.com/ergoapi/util/exjwt"

	"github.com/ergoapi/util/exgin"
	"github.com/gin-gonic/gin"
	"github.com/storyicon/grbac"
)

func init() {
	registerWithWeight("auth", 80, func() gin.HandlerFunc {
		return auth()
	})
}

func authrole(c *gin.Context) (roles []string, err error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if !strings.HasPrefix(bearerToken, "Bearer ") || len(strings.Fields(bearerToken)) != 2 {
		return nil, nil
	}
	token := strings.Fields(bearerToken)[1]

	claims, err := exjwt.Parse(token)
	if err != nil {
		return nil, err
	}
	uuid := claims["uuid"].(string)
	info := map[string]string{
		"username": claims["username"].(string),
		"uuid":     uuid,
	}
	c.Set("userinfo", info)
	return []string{uuid}, nil
}

// auth jwt auth
func auth() gin.HandlerFunc {
	rbacrule, err := grbac.New(grbac.WithRules(rbac.Rules()))
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Next()
			return
		}
		roles, err := authrole(c)
		if err != nil {
			exgin.GinsAbort(c, "token不合法")
			return
		}
		state, err := rbacrule.IsRequestGranted(c.Request, roles)
		if err != nil {
			exgin.GinsAbort(c, "权限校验失败")
			return
		}
		if !state.IsGranted() {
			exgin.GinsAbort(c, "暂无权限")
			return
		}
		c.Next()
	}
}
