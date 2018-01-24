package controller

import (
	"GoWeibo/src/logic"
	"fmt"

	"github.com/gin-gonic/gin"
)

// RegistController 注册
type RegistController struct{}

// RegisterRoute 注册路由
func (c RegistController) RegisterRoute(e *gin.Engine) {
	e.GET("/regist", UserRegist)
	e.POST("/login", Login)
}

// UserRegist 用户注册函数
func UserRegist(c *gin.Context) {

	userLogic := login.DefaultUser
	info, err := userLogic.CreateUser(c)
	fmt.Println(info)
	if err != nil {
		fail(c, err.Error())
		return
	}
	success(c, "注册成功")
}

// Login 用户登录
func Login(c *gin.Context) {
	userLogic := login.DefaultUser
	data, err := userLogic.UserLogin(c)
	if err != nil {
		fail(c, err.Error())
		return
	}
	success(c, data)
}
