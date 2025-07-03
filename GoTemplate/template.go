package GoTemplate

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http server start failed,errMsg:%v\n", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("GoTemplate file parse failed,errMsg:%v\n", err)
		return
	}
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	//渲染模板
	err = t.Execute(w, user)
	if err != nil {
		fmt.Printf("GoTemplate file render failed,errMsg:%v\n", err)
		return
	}
}
