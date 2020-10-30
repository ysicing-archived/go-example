// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package command

import (
	"app/models"
	gcron "app/pkg/cron"
	"app/pkg/gins"
	"app/pkg/middleware"
	"app/routers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ysicing/ext/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	taskscron := gcron.CronTasks{Cron: cron.New()}
	taskscron.Start()
	defer taskscron.Stop()
	addr := viper.GetString("server.listen")
	srv := &http.Server{
		Addr:    addr,
		Handler: gins.Gins,
	}
	if viper.GetBool("server.ssl.enable") {
		go startTls(gins.Gins)
	}
	go func() {
		logger.Slog.Infof("http listen to %v, pid is %v", addr, os.Getpid())
		//	utils.CheckAndExit(gins.Gins.Run(addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Slog.Fatal(err)
		}
	}()
	SetupGracefulStop(srv)
}

func startTls(e *gin.Engine) {
	tlsaddr := viper.GetString("server.ssl.listen")
	srv := &http.Server{
		Addr:    tlsaddr,
		Handler: e,
	}
	tlscfile := viper.GetString("server.ssl.cert")
	tlskfile := viper.GetString("server.ssl.key")
	go func() {
		logger.Slog.Infof("tls listen to %v, pid is %v", tlsaddr, os.Getpid())

		if err := srv.ListenAndServeTLS(tlscfile, tlskfile); err != nil && err != http.ErrServerClosed {
			logger.Slog.Fatal(err)
		}
	}()
	SetupGracefulStop(srv)
}

func SetupGracefulStop(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Slog.Info("Shutdown Server ...")
	Shutdown(srv)
}

func Shutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Slog.Info("[http server shutdown err:]", err)
	}

	select {
	case <-ctx.Done():
		logger.Slog.Info("[http server exit timeout of 5 seconds.]")
	default:

	}
	logger.Slog.Info("[http server exited.]")
}
