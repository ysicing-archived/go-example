// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package middleware

import (
	"app/pkg/jwt"
	"app/pkg/rbac"
	"github.com/gin-gonic/gin"
	"github.com/storyicon/grbac"
	"github.com/ysicing/ext/e"
	"strings"
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

	claims, err := jwt.JwtParse(token)
	if err != nil {
		return nil, err
	}
	role := claims["role"].(string)
	info := map[string]string{
		"username": claims["username"].(string),
		"role":     role,
	}
	c.Set("userinfo", info)
	return []string{role}, nil
}

// auth jwt auth
func auth() gin.HandlerFunc {
	rbacrule, err := grbac.New(grbac.WithRules(rbac.Rules()))
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		roles, err := authrole(c)
		if err != nil {
			c.JSON(200, e.Error(10403, "token不合法"))
			c.Abort()
			return
		}
		state, err := rbacrule.IsRequestGranted(c.Request, roles)
		if err != nil {
			c.JSON(200, e.Error(10400, "权限校验失败"))
			c.Abort()
			return
		}
		if !state.IsGranted() {
			c.JSON(200, e.Error(10401, ""))
			c.Abort()
			return
		}
		c.Next()
	}
}
