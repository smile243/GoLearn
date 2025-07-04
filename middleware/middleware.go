package middleware

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义中间件 函数类型必须是gin.HandlerFunc
func LoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 调用后续的函数 阻止是Abort
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print("方法执行耗时:" + latency.String())
		// 获取发送的 status
		status := c.Writer.Status()
		log.Println("请求状态:" + strconv.Itoa(status))
	}
}

func M1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("m1 in...")
	}
}
