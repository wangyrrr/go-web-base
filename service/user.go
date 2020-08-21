package service

import (
	"go-web-base/core/entity"
	"go-web-base/core/exception"
	"go-web-base/global"
)

func GetUserByUsernameAndPassword(username string, password string) (entity.User, error) {
	var user entity.User
	db := global.DB.Model(&user)
	err := db.Where("username = ? and password = ?", username, password).First(&user).Error
	if err != nil {
		return user, exception.NewServiceException("A0001", "用户错误")
	} else {
		return user, err
	}
}
