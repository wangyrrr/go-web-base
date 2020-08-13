package v1

import (
	"github.com/gin-gonic/gin"
	"go-web-base/core/bo"
	"go-web-base/global/response"
	"go-web-base/service"
)

func GetDemo(c *gin.Context) {
	var pageInfo bo.PageInfo
	c.ShouldBindQuery(&pageInfo)
	result := service.GetDemo(pageInfo)
	response.SuccessWithData(result, c)
}
