package initializ

import (
	"airad/common/support"
	"fmt"
	"github.com/astaxie/beego"
)

func InitDatabase() {
	if beego.AppConfig.String("runmode") == "dev" {
		support.IsDebugSql = true
	}
	var err error
	err = support.GetMysqlConnInstance().InitDataPool("airad")
	if err != nil {
		fmt.Println("init db airad error")
	}

	err = support.GetMysqlConnInstance().InitDataPool("life")
	if err != nil {
		fmt.Println("init db airad error")
	}
}
