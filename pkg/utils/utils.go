// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package utils

import (
	"github.com/ysicing/ext/logger/zlog"
)

// CheckAndExit check & exit
func CheckAndExit(err error) {
	if err != nil {
		zlog.Fatal("err: %v", err)
	}
}
