package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func success(c *gin.Context, data interface{}) error {
	result := map[string]interface{}{
		"ok":   1,
		"msg":  "操作成功",
		"data": data,
	}

	_, err := json.Marshal(result)
	if err != nil {
		return err
	}

	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "操作成功",
		"data": data,
	})
	return nil
}

func fail(ctx *gin.Context, msg string, codes ...int) error {
	code := 1
	if len(codes) > 0 {
		code = codes[0]
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"code": code,
	})
	fmt.Println("error : ", msg)

	return nil
}
