// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package server

import (
	"app/models"
	"app/pkg/cron"
	"app/pkg/gins"
	"app/pkg/middleware"
	"app/routers"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/ergoapi/zlog"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve(ctx context.Context) error {
	models.Init()
	gins.GinInit()
	middleware.Init()
	routers.Init()
	defer cron.Cron.Stop()
	cron.Cron.Start()
	addr := viper.GetString("server.listen")
	srv := &http.Server{
		Addr:    addr,
		Handler: gins.Gins,
	}
	if viper.GetBool("server.ssl.enable") {
		go startTLS(ctx, gins.Gins)
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			zlog.Error("Failed to stop http server, error: %s", err)
		}
		zlog.Info("server exited.")
	}()
	zlog.Info("http listen to %v, pid is %v", addr, os.Getpid())
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zlog.Error("Failed to start http server, error: %s", err)
		return err
	}
	return nil
}

func startTLS(ctx context.Context, e *gin.Engine) {
	tlsaddr := viper.GetString("server.ssl.listen")
	srv := &http.Server{
		Addr:    tlsaddr,
		Handler: e,
	}
	tlscfile := viper.GetString("server.ssl.cert")
	tlskfile := viper.GetString("server.ssl.key")
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			zlog.Error("failed to stop tls server, error: %s", err)
		}
		zlog.Info("tls server exited.")
	}()
	zlog.Info("tls listen to %v, pid is %v", tlsaddr, os.Getpid())
	if err := srv.ListenAndServeTLS(tlscfile, tlskfile); err != nil && err != http.ErrServerClosed {
		zlog.Fatal("failed to start tls server, error: %s", err)
	}
}
