// Copyright (c) 2022 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/log/glog"
	"github.com/ergoapi/util/ztime"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/guregu/null.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

var GDB *gorm.DB
var Migrates []interface{}

type Model struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id"`
	CreatedAt null.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt null.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt null.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func Migrate(obj interface{}) {
	Migrates = append(Migrates, obj)
}

func Init() {
	var err error
	dbtype := viper.GetString("db.type")
	dbdsn := viper.GetString("db.dsn")
	switch dbtype {
	case "mysql":
		GDB, err = gorm.Open(mysql.Open(dbdsn), &gorm.Config{
			Logger: &glog.DefaultGLogger,
		})
	default:
		GDB, err = gorm.Open(sqlite.Open(dbdsn), &gorm.Config{
			Logger: &glog.DefaultGLogger,
		})
	}
	if err != nil {
		logrus.Panicf("setup db err: %v", err.Error())
	}
	if viper.GetBool("db.metrics.enable") {
		dbname := viper.GetString("db.metrics.name")
		if len(dbname) == 0 {
			dbname = "example" + ztime.GetToday()
		}
		GDB.Use(prometheus.New(prometheus.Config{
			DBName: dbname,
			//RefreshInterval:  0,
			//PushAddr:         "",
			//StartServer:      false,
			//HTTPServerPort:   0,
			//MetricsCollector: nil,
		}))
	}
	dbcfg, _ := GDB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbcfg.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	dbcfg.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	dbcfg.SetConnMaxLifetime(time.Hour)

	if err := GDB.AutoMigrate(Migrates...); err != nil {
		logrus.Errorf("auto migrate table err: %v", err.Error())
	}
	logrus.Info(color.SGreen("create db engine success..."))

	InitSalt()
	InitAdmin()
}
