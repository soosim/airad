package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/rs/xid"
)

func InitFilter() {
	beego.InsertFilter("/*", beego.BeforeRouter, mainFilter)
}

/**
主过滤器
 */
var mainFilter = func(ctx *context.Context) {
	ctx.Input.SetData("commonLogId", xid.New().String())
}
