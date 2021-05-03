package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	//打印日志到控制台
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	r := gin.Default()
	//全局拦截器
	r.Use(GlobalMiddleware)
	//匹配路径的拦截器
	r.GET("/path",AuthMiddleWare(), func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"data":"OK"})
		fmt.Println("path")
	})
	r.Run(":8000")
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		authorized := check(token)
		if authorized {
			context.Next()
			return
		}
		context.JSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized"})
		context.Abort()
		return
	}
}

func check(token string) bool {
	if token == "ginAuth" {
		return true
	}else {
		return false
	}
}

//全局中间件 允许跨域
func GlobalMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	fmt.Println("global")
	c.Next()
}
