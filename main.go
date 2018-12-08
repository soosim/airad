package main

import (
	"airad/common/filter"
	"airad/common/initializ"
	"airad/module/common"
	_ "airad/routers"

	"github.com/astaxie/beego"
)

func main() {
	filter.InitFilter()
	initializ.InitSql()
	// utils.InitTemplate()
	initializ.InitCache()
	initializ.InitBootStrap()
	beego.ErrorController(&common.ErrorController{})

	beego.Run()
}
