// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

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
