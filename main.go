// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"app/cmd"
	"github.com/ergoapi/errors"
	"runtime"
)

// @title Go Example API
// @version 0.0.3
// @description This is a sample server Petstore server.

// @contact.name ysicing
// @contact.url http://github.com/ysicing
// @contact.email i@ysicing.me

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

func main() {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	errors.CheckAndExit(cmd.Execute())
}
