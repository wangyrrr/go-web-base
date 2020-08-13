package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-web-base/global"
	"os"
)

func InitMysql() {
	mysqlConfig := global.CONFIG.Mysql
	db, err := gorm.Open("mysql", mysqlConfig.Username + ":" + mysqlConfig.Password + "@" + mysqlConfig.Url)
	if err != nil {
		global.LOG.Error("mysql连接错误:", err)
		os.Exit(0)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tb_" + defaultTableName
	}
	db.DB().SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	db.LogMode(mysqlConfig.LogMode)
	global.DB = db
}