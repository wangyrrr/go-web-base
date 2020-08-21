package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const SUCCESS = 200
const ERROR = 500

type Response struct {
	Code int         `json: "code"`
	Msg  string      `json:"msg"`
	//Tip  string      `json:"tip"`
	Data interface{} `json:"data"`
}

/**
通用返回结果
*/
func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{code, msg, data})
}

func Success(c *gin.Context) {
	Result(SUCCESS, "SUCCESS", nil, c)
}

func SuccessWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, "SUCCESS", data, c)
}

func Exception(c *gin.Context) {
	Result(ERROR, "系统错误", nil, c)
}

func ExceptionWithMsg(msg string, c *gin.Context) {
	Result(ERROR, msg, nil, c)
}

func ExceptionWithDetail(msg string, data interface{}, c *gin.Context) {
	Result(ERROR, msg, data, c)
}
