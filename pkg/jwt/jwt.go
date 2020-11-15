// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

var jwtSecret []byte

func JwtAuth(username string) (t string, err error) {
	var expSecond int
	now := time.Now()
	expSecond = viper.GetInt("jwt.expsecond")
	// default token exp time is 86400s 60 * 60 * 24
	if expSecond == 0 {
		expSecond = 86400
	}
	if username == "admin" {
		expSecond = 86400000 // 1000d
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
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
