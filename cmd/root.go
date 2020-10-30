// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"app/cmd/command"
	"app/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/exmisc"
	"github.com/ysicing/ext/utils/exos"
)

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "go example by ysicing",
	Long: `Go example by ysicing.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Slog.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&command.CfgFile, "config", "", "config file (default is /conf/example.yml)")
	rootCmd.PersistentFlags().BoolVar(&command.Debug, "debug", false, "enable debug logging")
	rootCmd.AddCommand(command.NewVersionCommand(), command.ServerCommand())
	logcfg := &logger.Config{Simple: true, ConsoleOnly: false}
	logger.InitLogger(logcfg)
}

func initConfig() {
	if command.CfgFile == "" {
		command.CfgFile = command.Defaultcfgpath
		if exos.IsMacOS() {
			command.CfgFile = "./example.yaml"
		}
	}
	viper.SetConfigFile(command.CfgFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		utils.ShowDebugMsg("Using config file:", exmisc.SGreen(viper.ConfigFileUsed()))
	}
	// reload
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		utils.ShowDebugMsg("config changed: ", exmisc.SGreen(in.Name))
	})
	if command.Debug {
		viper.Set("server.debug", true)
	}
}
