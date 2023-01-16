// Copyright (c) 2023 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package cron

import (
	"app/pkg/prom"

	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/zos"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var Cron *Client

type Client struct {
	client *cron.Cron
}

func New() *Client {
	return &Client{client: cron.New()}
}

func (c *Client) Start() {
	logrus.Info(color.SGreen("start cron tasks"))
	c.client.Start()
}

func (c *Client) Add(spec string, cmd func()) (int, error) {
	id, err := c.client.AddFunc(spec, cmd)
	if err != nil {
		return 0, err
	}
	logrus.Infof("add cron: %v", id)
	return int(id), nil
}

func (c *Client) Remove(id int) {
	c.client.Remove(cron.EntryID(id))
}

func (c *Client) Default() {
	logrus.Info("add default cron")
	id, err := c.Add("@every 30s", func() {
		logrus.Debug(zos.GetHostname())
		prom.CronRunTimesCounter.WithLabelValues("default_cron").Inc()
	})
	if err != nil {
		logrus.Errorf("add default cron error: %s", err)
		return
	}
	logrus.Infof("add default cron [%d] success", id)
}

func (c *Client) Stop() {
	logrus.Info(color.SGreen("stop cron tasks"))
	c.client.Stop()
}

func (c *Client) List() []cron.Entry {
	return c.client.Entries()
}

func init() {
	Cron = New()
}
