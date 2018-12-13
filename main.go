package main

import (
	"airad/common/filter"
	"airad/common/initializ"
	"airad/module/common"
	_ "airad/routers"
	"runtime"

	"github.com/astaxie/beego"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	filter.InitFilter()
	initializ.InitDatabase()
	// utils.InitTemplate()
	initializ.InitCache()
	initializ.InitBootStrap()
	beego.ErrorController(&common.ErrorController{})

	beego.Run()
}
