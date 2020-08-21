package initialiaze

import (
	"go-web-base/core/entity"
	"go-web-base/global"
)

func DbTables() {
	db := global.DB
	db.AutoMigrate(entity.Demo{},
		entity.User{})
	global.LOG.Info("初始化数据库表成功")
}
