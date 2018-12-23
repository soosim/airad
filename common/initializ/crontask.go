package initializ

import (
	"airad/common/support"
	"airad/module/life/task"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
	"time"
)

var cronLogPrefix = "[cronTask] "
var taskList = make(map[string]*support.CronTask)

func taskInit() {
	taskList["test"] = &support.CronTask{Name: "test", Spec: "*/10 * * * * *", WorkFunc: func() {
		fmt.Println("Every 10 second")
	}}

	taskList["test1"] = &support.CronTask{Name: "test1", Spec: "*/30 * * * * *", WorkFunc: func() {
		fmt.Println("Every 30 second")
	}}

	taskList["lifeEveryTenSecond"] = &support.CronTask{Name: "lifeEveryTenSecond", Spec: "*/10 * * * * *", WorkFunc: task.EveryTenSecond}
}

func InitCronTask() {
	taskInit()
	if 0 == len(taskList) {
		logs.Info(cronLogPrefix + "no crontab to run")
		return
	}

	c := cron.NewWithLocation(time.Local)
	for _, v := range taskList {
		logs.Info(cronLogPrefix + "add cron task :" + v.Name)
		c.AddFunc(v.Spec, v.WorkFunc)
	}
	c.Start()
}
