// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package middleware

import (
	"app/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// JWTAuth jwt auth
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(bearerToken, "Bearer ") || len(strings.Fields(bearerToken)) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":   "Invalid  Token",
				"timestamp": time.Now().Unix(),
			})
			c.Abort()
			return
		}
		token := strings.Fields(bearerToken)[1]

		claims, err := jwt.JwtParse(token)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message":   err.Error(),
				"timestamp": time.Now().Unix(),
			})
			c.Abort()
			return
		}
		info := map[string]string{
			"username":   claims["username"].(string),
		}
		// c.SetCookie("username", claims["username"].(string), viper.GetInt("server.cookie"), "/", ".ysicing.local", true, true)
		c.Set("userinfo", info)
		c.Next()
	}
}