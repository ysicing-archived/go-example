// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cron

import (
	"app/pkg/prom"
	"github.com/robfig/cron/v3"
	"github.com/ysicing/ext/logger/zlog"
	"github.com/ysicing/ext/misc"
	"github.com/ysicing/ext/zos"
)

type CronTasks struct {
	Cron *cron.Cron
}

func (c *CronTasks) Start() {
	zlog.Info(misc.SGreen("start cron tasks"))
	c.Cron.AddFunc("@every 30s", func() {
		zlog.Debug(zos.GetHostname())
		prom.CronRunTimesCounter.WithLabelValues("default_cron").Inc()
	})
	c.Cron.Start()
}

func (c *CronTasks) Stop() {
	zlog.Info(misc.SGreen("stop cron tasks"))
	defer c.Cron.Stop()
}
