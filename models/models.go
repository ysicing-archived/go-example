// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package models

import (
	"github.com/spf13/viper"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/exmisc"
	"gopkg.in/guregu/null.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	dblog "gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
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
	newLogger := dblog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		dblog.Config{
			SlowThreshold: time.Second,
			Colorful:      false,
			LogLevel:      dblevel(),
		})
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
	if err != nil {
		logger.Slog.Exitf(-1, "setup db err: %v", err.Error())
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

func dblevel() dblog.LogLevel {
	switch strings.ToLower(viper.GetString("db.loglevel")) {
	case "silent":
		return dblog.Silent
	case "info":
		return dblog.Info
	case "error":
		return dblog.Error
	case "warn":
		return dblog.Warn
	default:
		return dblog.Silent
	}
}
