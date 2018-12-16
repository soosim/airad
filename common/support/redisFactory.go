package support

import (
    "github.com/astaxie/beego"
    "github.com/go-redis/redis"
    "sync"
)

var redisClient *redis.Client
var onceRedis sync.Once

func InitRedisClient() {
    redisClient = redis.NewClient(&redis.Options{
        Addr:     beego.AppConfig.String("redis.host"),
        Password: beego.AppConfig.String("redis.password"),
        DB:       0, // use default DB
    })
}

func GetRedisClient() *redis.Client {
    return redisClient
}
