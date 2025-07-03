package form

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	//获取form提交的参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
	})
}
