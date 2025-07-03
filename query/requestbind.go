package query

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequestBind(c *gin.Context) {
	var msg Message
	//参数绑定 既能绑定到get 也能绑定到post上
	err := c.ShouldBind(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Printf("%v\n", msg)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
