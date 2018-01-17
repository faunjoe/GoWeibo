package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Regist 注册
type Regist struct{}

// RegisterRoute 注册路由
func (c Regist) RegisterRoute(e *gin.Engine) {
	e.GET("/regist", userRegist)
}

func userRegist(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	fmt.Println(email)
	fmt.Println(password)
	success(c, "userRegist")
}
