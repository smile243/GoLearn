package query

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	//struct tag用法 因为前端不一定首字母大写
	Content string `json:"content"`
	Age     int
}

func RequestByMap(c *gin.Context) {
	//map类型
	data := gin.H{
		"message": "hello yjl",
	}
	c.JSON(http.StatusOK, data)
}

func RequestByStruct(c *gin.Context) {
	//结构体类型
	data2 := Message{
		Content: " yjl",
		Age:     20,
	}
	c.JSON(http.StatusOK, data2)
}
