package db

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	// 初始化数据库
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MasterDB 主数据库
var MasterDB *gorm.DB

func init() {
	// 启动应用就打开数据库连接
	if err := initEngine(); err != nil {
		fmt.Println("数据库没有打开")
		panic(err)
	}
}

func initEngine() error {
	var err error
	MasterDB, err = gorm.Open("mysql", "root:123456@/weibo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	MasterDB.LogMode(true)
	return nil
}

// StdMasterDB 获取数据库
func StdMasterDB() *sql.DB {
	return MasterDB.DB()
}
