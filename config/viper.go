package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-web-base/global"
	"os"
)

const defaultConfigFileName = "application"
const configType = "yaml"

func InitViper() {
	v := viper.New()
	v.SetConfigType(configType)
	v.AddConfigPath("./resource/")
	v.SetConfigName(defaultConfigFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	configs := v.AllSettings()
	for k, val := range configs {
		v.SetDefault(k, val)
	}
	//global.LOG.Info("默认配置：", configs)
	env := os.Getenv("env")
	if env != "" {
		v.SetConfigName(defaultConfigFileName + "-" + env)
		err = v.ReadInConfig()
		if err != nil {
			fmt.Println("没有找到指定环境配置文件", err)
		}
	}
	err = v.Unmarshal(&global.CONFIG)
	if err != nil {
		panic(fmt.Errorf("配置映射错误: %s \n", err))
	}
	//global.LOG.Infof("%s环境配置：%s", env, global.CONFIG)
	global.VIPER = v
}
