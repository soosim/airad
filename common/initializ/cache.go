package initializ

import (
	"airad/common/support"
	"github.com/astaxie/beego/logs"

	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
)

func InitCache() {
	if err := support.InitRedisClient(); nil != err {
		logs.Error("init redis error :", err)
	}
}
