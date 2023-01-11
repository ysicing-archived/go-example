// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

func Auth(username string, role ...string) (t string, err error) {
	var expSecond int64
	now := time.Now()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username

	if len(role) > 0 && strings.HasPrefix(role[0], "ptv") {
		claims["role"] = role[0]
		switch role[0] {
		case "ptv6":
			expSecond = 86400000 // 1000d
		case "ptv4":
			expSecond = 86400000 // 1000d
		case "ptv2":
			expSecond = 86400 // 1d 86400s 60 * 60 * 24
		default:
			expSecond = 600 // 10m
		}
	} else {
		claims["role"] = "ptv0"
		expSecond = 60 // 1m
	}

	claims["exp"] = now.Add(time.Duration(expSecond) * time.Second).Unix()
	t, err = token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("JWT Generate Failure")
	}
	return t, nil
}

func Parse(tokenstring string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Token invalid")
	}
	claim := token.Claims.(jwt.MapClaims)
	return claim, nil
}
