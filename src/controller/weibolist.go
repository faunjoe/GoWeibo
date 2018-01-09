package controller

import (
	"github.com/gin-gonic/gin"
)

// WeiboListController 微博列表控制器
type WeiboListController struct{}

// RegisterRoute 注册路由
func (c WeiboListController) RegisterRoute(e *gin.Engine) {
	e.GET("/list", test)
}

func (c *WeiboListController) list(e *gin.Engine) string {
	return ""
}

func test(c *gin.Context) {
	success(c, "ddd")
}
