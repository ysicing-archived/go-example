// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package models

import (
	"github.com/spf13/viper"
	"github.com/ysicing/ext/exlog/dblog"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/exmisc"
	"github.com/ysicing/ext/utils/extime"
	"gopkg.in/guregu/null.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
	"time"
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
	dbmode := viper.GetBool("db.debug")
	newLogger := dblog.New(logger.Slog, dbmode)
	switch dbtype {
	case "mysql":
		GDB, err = gorm.Open(mysql.Open(dbdsn), &gorm.Config{
			Logger: newLogger,
		})
	default:
		GDB, err = gorm.Open(sqlite.Open(dbdsn), &gorm.Config{
			Logger: newLogger,
		})
	}
	if viper.GetBool("db.metrics.enable") {
		dbname := viper.GetString("db.metrics.name")
		if len(dbname) == 0 {
			dbname = "example" + extime.GetToday()
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
	if err != nil {
		logger.Slog.Fatalf("setup db err: %v", err.Error())
	}
	dbcfg, _ := GDB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbcfg.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	dbcfg.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	dbcfg.SetConnMaxLifetime(time.Hour)

	if err := GDB.AutoMigrate(Migrates...); err != nil {
		logger.Slog.Errorf("auto migrate table err: %v", err.Error())
	}
	logger.Slog.Info(exmisc.SGreen("create db engine success..."))
}
