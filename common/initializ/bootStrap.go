package initializ

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"os/signal"
	"syscall"
)

func InitBootStrap() {
	graceful, _ := beego.AppConfig.Bool("Graceful")
	if !graceful {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go handleSignals(sigs)
	}
	beego.SetLogger(logs.AdapterFile, `{"filename":"public/logs/logs.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	// beego.SetLogger(logs.AdapterFile, `{"filename":"logs/logs.log"}`)
	// logs.EnableFuncCallDepth(true)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.SetStaticPath("/swagger", "swagger")
		beego.SetLevel(beego.LevelDebug)
	} else if beego.BConfig.RunMode == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}

	// 模板目录
	beego.SetViewsPath("public/views")
}

func handleSignals(c chan os.Signal) {
	switch <-c {
	case syscall.SIGINT, syscall.SIGTERM:
		fmt.Println("Shutdown quickly, bye...")
	case syscall.SIGQUIT:
		fmt.Println("Shutdown gracefully, bye...")
		// do graceful shutdown
	}
	os.Exit(0)
}
