package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-web-base/api/v1"
)

func InitUserRouter(router *gin.RouterGroup) {
	apiRouter := router.Group("/v1/user")
	apiRouter.POST("/login", v1.Login)
}
