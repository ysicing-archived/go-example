// Copyright (c) 2023 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package main

import (
	"app/cmd"
	"runtime"

	errors "github.com/ergoapi/util/exerror"
)

// @title Go Example API
// @version 0.1.0
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
