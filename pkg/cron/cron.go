// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/exmisc"
	"github.com/ysicing/ext/utils/exos"
)

type CronTasks struct {
	Cron *cron.Cron
}

func (c *CronTasks) Start() {
	logger.Slog.Info(exmisc.SGreen("start cron tasks"))
	c.Cron.AddFunc("@every 1s", func() {
		logger.Slog.Info(exos.GetHostname())
	})
	c.Cron.Start()
}

func (c *CronTasks) Stop() {
	logger.Slog.Info(exmisc.SGreen("stop cron tasks"))
	defer c.Cron.Stop()
}
