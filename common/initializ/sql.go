package initializ

import (
	"airad/common/support"
)

func InitSql() {
	/*if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	initAiradDB()
	initLifeDB()*/
	support.GetMysqlConnInstance().InitDataPool()
}

/*func initAiradDB() {
	dbUser := beego.AppConfig.String("airad.db.user")
	dbPassword := beego.AppConfig.String("airad.db.password")
	dbHost := beego.AppConfig.String("airad.db.host")
	dbPort := beego.AppConfig.String("airad.db.port")
	dbName := beego.AppConfig.String("airad.db.name")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if nil != err {
		fmt.Println("init database life error. the application will be shutdown")
		os.Exit(1)
	}

	// 自动建表功能关闭
	/*if err := orm.RunSyncdb("default", false, true); err != nil {
		fmt.Println(err)
	}*/
	//orm.RunCommand()
//}

/*func initLifeDB() {
	dbUser := beego.AppConfig.String("life.db.user")
	dbPassword := beego.AppConfig.String("life.db.password")
	dbHost := beego.AppConfig.String("life.db.host")
	dbPort := beego.AppConfig.String("life.db.port")
	dbName := beego.AppConfig.String("life.db.name")
	maxIdle := 5
	maxConn := 5
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err := orm.RegisterDataBase("life", "mysql", dsn, maxIdle, maxConn); err != nil {
		fmt.Println("init database life error. the application will be shutdown")
		os.Exit(1)
	}
}*/
