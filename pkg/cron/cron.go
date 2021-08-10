// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cron

import (
	"app/pkg/prom"
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/zos"
	"github.com/ergoapi/zlog"
	"github.com/robfig/cron/v3"
)

type CronTasks struct {
	Cron *cron.Cron
}

func (c *CronTasks) Start() {
	zlog.Info(color.SGreen("start cron tasks"))
	c.Cron.AddFunc("@every 30s", func() {
		zlog.Debug(zos.GetHostname())
		prom.CronRunTimesCounter.WithLabelValues("default_cron").Inc()
	})
	c.Cron.Start()
}

func (c *CronTasks) Stop() {
	zlog.Info(color.SGreen("stop cron tasks"))
	defer c.Cron.Stop()
}
