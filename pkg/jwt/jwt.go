// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"strings"
	"time"
)

var jwtSecret []byte

func JwtAuth(username string, role ...string) (t string, err error) {
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
		return "", errors.New("JWT Generate Failure")
	}
	return t, nil
}

func JwtParse(tokenstring string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("Token invalid")
	}
	claim := token.Claims.(jwt.MapClaims)
	return claim, nil
}
