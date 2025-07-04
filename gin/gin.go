package gin

import (
	"hello/form"
	"hello/query"
	"hello/upload"

	"github.com/gin-gonic/gin"
)

func jsonDemo(router *gin.Engine) {
	//路由分组
	userGroup := router.Group("/json")
	{
		userGroup.GET("/requestByMap", query.RequestByMap)
		userGroup.GET("/requestByStruct", query.RequestByStruct)
	}

}

func queryStringDemo(router *gin.Engine) {
	router.GET("json/requestWithString", query.RequestWithString)
}

func formDemo(router *gin.Engine) {
	router.GET("/login", form.GetLogin)
	router.POST("/login", form.Login)
}

func requestBodyDemo(router *gin.Engine) {
	router.POST("/requestBody", query.RequestBind)
}

func uploadDemo(router *gin.Engine) {
	router.GET("/upload", upload.GetUpload)
	router.POST("/upload", upload.Upload)
}

func Gin() {
	router := gin.Default()
	//获取所有得用LoadHTMLGlob 不可用LoadHTMLFiles
	router.LoadHTMLGlob("template/*")
	jsonDemo(router)
	queryStringDemo(router)
	formDemo(router)
	requestBodyDemo(router)
	uploadDemo(router)
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
