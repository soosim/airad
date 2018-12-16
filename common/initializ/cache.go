package initializ

import (
    "airad/common/support"
    "fmt"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/cache"
    _ "github.com/astaxie/beego/cache/memcache"
    _ "github.com/astaxie/beego/cache/redis"
)

func InitCache() {
    // cacheConfig := beego.AppConfig.String("cache")
    support.Cc = nil
    support.InitRedisClient()

    /*if "redis" == cacheConfig {
        initRedis()
    } else {
        initMemcache()
    }*/

    //fmt.Println("[cache] use", cacheConfig)
}

func initMemcache() {
    var err error
    support.Cc, err = cache.NewCache("memcache", `{"conn":"`+beego.AppConfig.String("memcache_host")+`"}`)

    if err != nil {
        beego.Info(err)
    }

}

func initRedis() {
    var err error

    defer func() {
        if r := recover(); r != nil {
            fmt.Println("initial redis error caught: %v\n", r)
            support.Cc = nil
        }
    }()

    support.Cc, err = cache.NewCache("redis", `{"conn":"`+beego.AppConfig.String("redis.host")+`", 
			"password":"`+beego.AppConfig.String("redis.password")+`"}`)

    if err != nil {
        fmt.Println(err)
    }
}
