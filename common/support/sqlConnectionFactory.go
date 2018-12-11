package support

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type MysqlConnectionPool struct {
}

var instance *MysqlConnectionPool
var once sync.Once

var airadDb *gorm.DB
var lifeDb *gorm.DB
var err_db error

func GetMysqlConnInstance() *MysqlConnectionPool {
	once.Do(func() {
		instance = &MysqlConnectionPool{}
	})
	return instance
}

func (m *MysqlConnectionPool) InitDataPool(database string) (isSuccess bool) {
	dbUser := beego.AppConfig.String(database + ".db.user")
	dbPassword := beego.AppConfig.String(database + ".db.password")
	dbHost := beego.AppConfig.String(database + ".db.host")
	dbPort := beego.AppConfig.String(database + ".db.port")
	dbName := beego.AppConfig.String(database + ".db.name")
	if "airad" == database {
		airadDb, err_db = gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	} else if "airad" == database {
		lifeDb, err_db = gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	}

	if err_db != nil {
		fmt.Println("init " + database + " DB error")
		return false
	}
	return true
}

func (m *MysqlConnectionPool) GetAiradDB() (db_conn *gorm.DB) {
	return airadDb
}

func (m *MysqlConnectionPool) GetLifeDB() (db_conn *gorm.DB) {
	return lifeDb
}
