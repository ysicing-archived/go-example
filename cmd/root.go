// Copyright (c) 2023 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package cmd

import (
	"app/cmd/command"
	"app/constants"

	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/zos"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "example",
		Short: "go example by ysicing",
		Long: `Go example by ysicing.
`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(false)
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&constants.CfgFile, "config", "", "config file (default is /conf/example.yml)")
	rootCmd.PersistentFlags().BoolVar(&constants.Debug, "debug", false, "enable debug logging")
	rootCmd.AddCommand(command.NewVersionCommand(), command.ServerCommand())
}

func initConfig() {
	if constants.CfgFile == "" {
		constants.CfgFile = constants.Defaultcfgpath
		if zos.IsMacOS() {
			constants.CfgFile = "./example.yaml"
		}
	}
	viper.SetConfigFile(constants.CfgFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		logrus.Debugf("Using config file: %v", color.SGreen(viper.ConfigFileUsed()))
	}
	// reload
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Debugf("config changed: %v", color.SGreen(in.Name))
	})
	if constants.Debug {
		logrus.SetReportCaller(true)
		viper.Set("server.debug", true)
	}
}
