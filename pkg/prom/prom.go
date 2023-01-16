// Copyright (c) 2023 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	ns = "app"
)

var (
	CronRunTimesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: ns,
		Name:      "cron_run_time",
		Help:      "cron执行次数",
	}, []string{"name"})
)

func init() {
	prometheus.MustRegister(CronRunTimesCounter)
}
