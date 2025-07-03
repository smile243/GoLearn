package GoTemplate

import (
	"fmt"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("GoTemplate/base.tmpl", "GoTemplate/index.tmpl")
	if err != nil {
		fmt.Printf("parse GoTemplate failed,err:%v\n", err)
	}
	msg := "这是index页面"
	t.ExecuteTemplate(w, "index.tmpl", msg)
}

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("GoTemplate/base.tmpl", "GoTemplate/home.tmpl")
	if err != nil {
		fmt.Printf("parse GoTemplate failed,err:%v\n", err)
	}
	msg := "这是home页面"
	t.ExecuteTemplate(w, "home.tmpl", msg)
}
