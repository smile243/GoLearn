package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello yjl",
	})
}

func main() {
	router := gin.Default()
	router.GET("/hello", sayHello)
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}