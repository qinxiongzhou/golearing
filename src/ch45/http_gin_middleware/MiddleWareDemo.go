package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/path",AuthMiddleWare(), func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"data":"OK"})
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
