package GoTemplate

import (
	"fmt"
	"net/http"
	"text/template"
)



func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http server start failed,errMsg:%v\n", err)
		return
	}
}

func index(w http.ResponseWriter,r *http.Request)  {
	t, err := template.ParseFiles("./base.tmpl","./index.tmpl")
	if err!=nil{
		fmt.Printf("parse GoTemplate failed,err:%v\n",err)
	}
	msg :="这是index页面"
	t.ExecuteTemplate(w,"index.tmpl",msg)
}

func home(w http.ResponseWriter,r *http.Request)  {
	t, err := template.ParseFiles("./base.tmpl","./home.tmpl")
	if err!=nil{
		fmt.Printf("parse GoTemplate failed,err:%v\n",err)
	}
	msg :="这是home页面"
	t.ExecuteTemplate(w,"home.tmpl",msg)
}

