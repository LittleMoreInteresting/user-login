package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/trace"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
	Salt       string
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	LogConf logx.LogConf

	Telemetry trace.Config
}
