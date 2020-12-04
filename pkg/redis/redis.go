// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package redis

import (
	"github.com/spf13/viper"
	"github.com/ysicing/ext/redis"
)

var RCClient *redis.Client

func CacheInit() {
	rediscfg := redis.RedisConfig{
		Maxidle:     31,
		Maxactive:   31,
		IdleTimeout: 200,
		Host:        viper.GetString("cache.host"),
		Port:        viper.GetInt("cache.port"),
		Password:    viper.GetString("cache.pass"),
	}
	cfg := redis.Config{RedisCfg: &rediscfg}
	RCClient = redis.New(&cfg)
}
