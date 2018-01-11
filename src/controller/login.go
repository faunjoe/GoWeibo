package controller

import "github.com/gin-gonic/gin"

// Regist 注册
type Regist struct{}

// RegisterRoute 注册路由
func (c Regist) RegisterRoute(e *gin.Engine) {
	e.GET("/regist", userRegist)
}

func userRegist(c *gin.Context) {
	success(c, "userRegist")
}
