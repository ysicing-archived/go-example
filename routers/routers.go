// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package routers

import (
	"app/pkg/gins"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/ergoapi/errors"
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/zlog"
	"github.com/gin-gonic/gin"
)

type routerFunc struct {
	Name   string
	Weight int
	Func   func(router *gin.Engine)
}

type routerSlice []routerFunc

var userRouterOnce sync.Once
var routers routerSlice

func (r routerSlice) Len() int {
	return len(r)
}

func (r routerSlice) Less(i, j int) bool {
	return r[i].Weight > r[j].Weight
}

func (r routerSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// register new router with key name
// key name is used to eliminate duplicate routes
// key name not case sensitive
func register(name string, f func(router *gin.Engine)) {
	registerWithWeight(name, 50, f)
}

// registerWithWeight register new router with weight
func registerWithWeight(name string, weight int, f func(router *gin.Engine)) {
	if weight > 100 || weight < 0 {
		errors.CheckAndExit(fmt.Errorf("router weight must be >= 0 and <=100"))
	}

	for _, r := range routers {
		if strings.EqualFold(name, r.Name) {
			errors.CheckAndExit(fmt.Errorf("router [%s] already register", r.Name))
		}
	}

	routers = append(routers, routerFunc{
		Name:   name,
		Weight: weight,
		Func:   f,
	})
}

// Init framework init
func Init() {
	userRouterOnce.Do(func() {
		sort.Sort(routers)
		for _, r := range routers {
			r.Func(gins.Gins)
			zlog.Debug(color.SGreen("load router [%s] success...", r.Name))
		}
		zlog.Info(color.SGreen("load router success..."))
	})
}
