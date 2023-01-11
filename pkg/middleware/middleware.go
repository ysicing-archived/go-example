// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package middleware

import (
	"app/pkg/gins"
	"fmt"
	"sort"

	"github.com/ergoapi/util/color"
	errors "github.com/ergoapi/util/exerror"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type middleware struct {
	Name        string
	HandlerFunc func() gin.HandlerFunc
	Weight      int
}

type middlewareSlice []middleware

var mws middlewareSlice

func (m middlewareSlice) Len() int { return len(m) }

func (m middlewareSlice) Less(i, j int) bool { return m[i].Weight > m[j].Weight }

func (m middlewareSlice) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

// registering new middleware
func register(name string, handlerFunc func() gin.HandlerFunc) {
	mw := middleware{
		HandlerFunc: handlerFunc,
		Weight:      50,
		Name:        name,
	}
	mws = append(mws, mw)
}

// registering new middleware with weight
func registerWithWeight(name string, weight int, handlerFunc func() gin.HandlerFunc) {

	if weight > 100 || weight < 0 {
		errors.CheckAndExit(fmt.Errorf("middleware weight must be >= 0 and <=100"))
	}

	mw := middleware{
		HandlerFunc: handlerFunc,
		Weight:      weight,
		Name:        name,
	}
	mws = append(mws, mw)
}

// Init middleware init func
func Init() {
	sort.Sort(mws)
	for _, mw := range mws {
		gins.Gins.Use(mw.HandlerFunc())
		logrus.Debugf("load middleware: %v", mw.Name)
	}
	logrus.Info(color.SGreen("load middleware success..."))
}
