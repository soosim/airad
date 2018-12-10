package filter

import (
	"airad/common/base"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/rs/xid"
)

func InitFilter() {
	beego.InsertFilter("/*", beego.BeforeRouter, mainFilter)
	beego.InsertFilter("/v1/*", beego.BeforeRouter, loginFilter)
}

//主过滤器
var mainFilter = func(ctx *context.Context) {
	ctx.Input.SetData("commonLogId", xid.New().String())
}

// 登陆校验过滤器
var loginFilter = func(ctx *context.Context) {
	token := ctx.Input.Header("token")

	// TODO : token校验
	// TODO : 登陆用户信息保存context

	if "" == token {
		// ctx.Redirect(302, "/")
		ctx.Output.JSON(base.ErrExpired, false, false)
		return
	}
}
