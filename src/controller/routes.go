package controller

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册路由
func RegisterRoutes(e *gin.Engine) {
	// 在这里调用控制器注册路由
	new(WeiboListController).RegisterRoute(e)
	new(RegistController).RegisterRoute(e)
}
