package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-web-base/api/v1"
)

func InitDemoRouter(router *gin.RouterGroup) {
	apiRouter := router.Group("/v1/demo")
	apiRouter.GET("getDemo", v1.GetDemo)
}
