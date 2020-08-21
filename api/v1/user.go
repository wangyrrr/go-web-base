package v1

import (
	"github.com/gin-gonic/gin"
	"go-web-base/core/entity"
	"go-web-base/core/vo"
	"go-web-base/global/response"
	"go-web-base/service"
	"reflect"
)

func Login(c *gin.Context) {
	var user entity.User
	err := c.BindJSON(&user)
	if err != nil {
		// 绑定参数错误
		response.ExceptionWithMsg("参数传递错误", c)
	}
	user, err = service.GetUserByUsernameAndPassword(user.Username, user.Password)
	if err != nil {
		if reflect.TypeOf(err).Name() == "ServiceException" {
			response.ExceptionWithMsg(err.Error(), c)
		} else {
			response.ExceptionWithMsg(err.Error(), c)
		}
	} else {
		loginResult := vo.LoginResult{User: user,
			Token:        "xxxxxxxxxxx",
			ExpirationIn: 11111111111}
		response.SuccessWithData(loginResult, c)
	}
}
