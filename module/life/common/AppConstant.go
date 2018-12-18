package common

import (
	"airad/common/base"
)

const MODULE_NAME = "life"

var (
	LifeCachePrefix = base.CachePrefix + base.CacheSeparate + MODULE_NAME + base.CacheSeparate
)
