package base

import "github.com/astaxie/beego"

var (
	CacheSeparate = beego.AppConfig.DefaultString("cache::redis.separate", ":")
	CachePrefix   = beego.BConfig.AppName
)
