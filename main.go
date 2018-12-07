package main


import (
	_ "airad/routers"
	"airad/utils"


	"github.com/astaxie/beego"
	"airad/controllers"
)

func main() {
	utils.InitFilter()
	utils.InitSql()
	utils.InitTemplate()
	utils.InitCache()
	utils.InitBootStrap()
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}