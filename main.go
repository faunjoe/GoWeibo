package main

import (
	"GoWeibo/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 注册路由
	controller.RegisterRoutes(r)
	r.Run()
}
