// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package command

import (
	"app/models"
	"app/pkg/gins"
	"app/pkg/middleware"
	"app/pkg/utils"
	"app/routers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/logger"
	"os"
)

func ServerCommand() *cobra.Command {
	s := &cobra.Command{
		Use:   "server",
		Short: "core server",
		Run:   core,
	}
	return s
}

func core(cmd *cobra.Command, args []string) {
	models.Init()
	gins.GinInit()
	middleware.Init()
	routers.Init()
	addr := viper.GetString("server.listen")
	logger.Slog.Infof("listen to %v, pid is %v", addr, os.Getpid())
	utils.CheckAndExit(gins.Gins.Run(addr))
}
