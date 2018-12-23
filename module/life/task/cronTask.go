package task

import "github.com/astaxie/beego/logs"

func EveryTenSecond() {
	logs.Info("life module cron execing")
}
