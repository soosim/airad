package support

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type MysqlConnectionPool struct {
}

var instance *MysqlConnectionPool
var once sync.Once

var db *gorm.DB
var err_db error

func GetMysqlConnInstance() *MysqlConnectionPool {
	once.Do(func() {
		instance = &MysqlConnectionPool{}
	})
	return instance
}

func (m *MysqlConnectionPool) InitDataPool() (isSuccess bool) {
	db, err_db = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/airad?charset=utf8&parseTime=True&loc=Local")
	if err_db != nil {
		fmt.Println("init DB error")
		return false
	}
	return true
}

func (m *MysqlConnectionPool) GetMysqlDB() (db_conn *gorm.DB) {
	return db
}
