package support

import (
	"errors"
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

	poolStats := redisClient.PoolStats()
	if 0 == poolStats.TotalConns {
		return errors.New("pool TotalConns is zero")
	}
	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
