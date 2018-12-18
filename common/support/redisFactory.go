package support

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"sync"
)

var redisClient *redis.Client
var onceRedis sync.Once

func InitRedisClient() error {
	redisHost := beego.AppConfig.String("cache::redis.host")
	if "" == redisHost {
		panic("redisHost config error")
	}
	redisPass := beego.AppConfig.String("cache::redis.password")

	redisClient = redis.NewClient(&redis.Options{
		Addr:         redisHost,
		Password:     redisPass,
		DB:           0, // use default DB
		MinIdleConns: 1,
		PoolSize:     5,
	})

	_, err := redisClient.Ping().Result()
	if nil != err {
		return err
	}
	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
