package models

import (
	//"fmt"
	"github.com/astaxie/beego/config"
	//"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//返回带前缀的表名
func TableName(str string) string {
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	return appConf.String("database::") + str
}
