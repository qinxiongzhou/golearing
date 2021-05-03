package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"pong",
		})
	})

	/**
	冒号“：”加上一个参数名组成路由参数。可以使用context.Param()的方法读取其值
	可以匹配：
	/user/aoho
	/user/ryan
	不能匹配：
	/user
	/user/
	/user/aoho/
	*/
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK,"Hello %s",name)
	})

	/**
	gin除了提供冒号，还提供星号“*”处理参数
	可以匹配：
	/user/ryan/
	/user/ryan/send
	 */
	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK,message)
	})
	router.Run(":8000") // 默认监听0.0.0.0:8080
}
