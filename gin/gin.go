package gin

import (
	"hello/form"
	"hello/middleware"
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
	//默认使用logger,recovery中间件
	router := gin.Default()
	//获取所有得用LoadHTMLGlob 不可用LoadHTMLFiles
	router.LoadHTMLGlob("template/*")
	//全局注册自定义中间件
	router.Use(middleware.LoggerHandler(), middleware.M1())
	jsonDemo(router)
	queryStringDemo(router)
	formDemo(router)
	requestBodyDemo(router)
	uploadDemo(router)
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
