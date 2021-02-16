// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package routers

import (
	"app/pkg/gins"
	"app/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ext/logger/zlog"
	"github.com/ysicing/ext/misc"
	"sort"
	"strings"
	"sync"
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
		utils.CheckAndExit(fmt.Errorf("router weight must be >= 0 and <=100"))
	}

	for _, r := range routers {
		if strings.ToLower(name) == strings.ToLower(r.Name) {
			utils.CheckAndExit(fmt.Errorf("router [%s] already register", r.Name))
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
			zlog.Debug(misc.SGreen("load router [%s] success...", r.Name))
		}
		zlog.Info(misc.SGreen("load router success..."))
	})
}
