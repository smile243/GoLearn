package query

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestWithString(c *gin.Context) {
	//获取浏览器请求携带的参数 ?query=xxx
	//name := c.Query("query")
	//相当于map.getOrDefault
	name := c.DefaultQuery("query", "default_name")
	age := c.DefaultQuery("age", "18")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}
