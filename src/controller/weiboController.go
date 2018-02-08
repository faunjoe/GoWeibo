package controller

import (
	"github.com/gin-gonic/gin"
)

// WeiboController 微博列表控制器
type WeiboController struct{}

// RegisterRoute 注册路由
func (c WeiboController) RegisterRoute(e *gin.Engine) {
	e.GET("/addweibo", addWeibo)
}

func addWeibo(c *gin.Context) {

}
