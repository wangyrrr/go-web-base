package global

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go-web-base/core/bo"
	"go.uber.org/zap"
)

var (
	VIPER  *viper.Viper
	CONFIG bo.Server
	DB     *gorm.DB
	LOG    *zap.SugaredLogger
)

