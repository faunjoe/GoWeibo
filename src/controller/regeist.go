package controller

import (
	"GoWeibo/src/model"

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
	u := model.User{Email: email, Password: password}
	err := model.InsertUser(&u)
	if err != nil {
		fail(c, err.Error())
		return
	}
	success(c, "注册成功")
	// fmt.Println(email)
	// fmt.Println(password)
	// success(c, "userRegist")
}
