// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"app/cmd/command"
	"app/constants"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/logger/zlog"
	"github.com/ysicing/ext/misc"
	"github.com/ysicing/ext/zos"
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&constants.CfgFile, "config", "", "config file (default is /conf/example.yml)")
	rootCmd.PersistentFlags().BoolVar(&constants.Debug, "debug", false, "enable debug logging")
	rootCmd.AddCommand(command.NewVersionCommand(), command.ServerCommand())
	logcfg := &zlog.Config{Simple: true, WriteLog: false, WriteJSON: true, ServiceName: "example"}
	zlog.InitZlog(logcfg)
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
		zlog.Debug("Using config file: %v", misc.SGreen(viper.ConfigFileUsed()))
	}
	// reload
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zlog.Debug("config changed: %v", misc.SGreen(in.Name))
	})
	if constants.Debug {
		viper.Set("server.debug", true)
	}
}
