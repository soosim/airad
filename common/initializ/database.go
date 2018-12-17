package initializ

import (
	"airad/common/support"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func InitDatabase() {
	if beego.AppConfig.String("runmode") == "dev" {
		support.IsDebugSql = true
	}
	var err error
	err = support.GetMysqlConnInstance().InitDataPool("airad")
	if err != nil {
		logs.Error("init db airad error : ", err)
		fmt.Println("init db airad error : ", err)
	}

	err = support.GetMysqlConnInstance().InitDataPool("life")
	if err != nil {
		logs.Error("init db life error : ", err)
		fmt.Println("init db life error : ", err)
	}
}
