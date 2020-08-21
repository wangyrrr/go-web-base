package initialiaze

import (
	"github.com/gin-gonic/gin"
	"go-web-base/middleware"
	"go-web-base/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	// 处理跨域
	r.Use(middleware.Cors())
	// 统一路由管理
	apiGroup := r.Group("/api")
	router.InitDemoRouter(apiGroup)
	router.InitUserRouter(apiGroup)
	return r
}