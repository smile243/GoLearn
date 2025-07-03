package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func parse(c *gin.Context){
	//渲染模板
	c.HTML(http.StatusOK,"index.tmpl",gin.H{
		"title":"yjl",
		"网址":"<a href='https://liwenzhou.com'>李文周的博客</a>",
	})
}
func main()  {
	r := gin.Default()
	//自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	//加载静态文件
	r.Static("/cssStyle","../static")
	//解析模板
	r.LoadHTMLFiles("./index.tmpl")
	r.GET("/index",parse)
	r.Run(":9090")
}
