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

var dbConn = make(map[string]*gorm.DB)
var IsDebugSql bool

func init() {
}

func GetMysqlConnInstance() *MysqlConnectionPool {
	once.Do(func() {
		instance = &MysqlConnectionPool{}
	})
	return instance
}

func (m *MysqlConnectionPool) InitDataPool(database string) (errDb error) {
	dbUser := beego.AppConfig.String(database + ".db.user")
	dbPassword := beego.AppConfig.String(database + ".db.password")
	dbHost := beego.AppConfig.String(database + ".db.host")
	dbPort := beego.AppConfig.String(database + ".db.port")
	dbName := beego.AppConfig.String(database + ".db.name")
	dbCharset := beego.AppConfig.String(database + ".db.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset,
	)
	dbConn[database], errDb = gorm.Open("mysql", dsn)
	dbConn[database].DB().SetMaxIdleConns(10)
	dbConn[database].DB().SetMaxOpenConns(100)
	fmt.Println("init " + database + " Success")

	if errDb != nil {
		fmt.Println("init " + database + " DB error")
	}
	return errDb
}

func (m *MysqlConnectionPool) GetDBConn(database string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if _, ok := dbConn[database]; ok {
		db = dbConn[database]
		db.DB().Ping()
	} else {
		if err = GetMysqlConnInstance().InitDataPool(database); err == nil {
			db = dbConn[database]
		}
	}

	if IsDebugSql && db != nil {
		db = db.LogMode(true).Debug()
	}
	return db, err
}
