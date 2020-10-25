// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package cmd

import (
	"app/cmd/command"
	"app/pkg/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/utils/exmisc"
	"github.com/ysicing/ext/utils/exos"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "go example by ysicing",
	Long: `Go example by ysicing.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&command.CfgFile, "config", "", "config file (default is /conf/example.yml)")
	rootCmd.PersistentFlags().BoolVar(&command.Debug, "debug", false, "enable debug logging")
	rootCmd.AddCommand(command.NewVersionCommand(), command.ServerCommand())
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
