package GoTemplate

import (
	"fmt"
	"net/http"
	"text/template"
)


func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http server start failed,errMsg:%v\n", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	t := template.New("helloV3.tmpl")
	autochat := func(name string) (string, error) {
		return name + "你好啊", nil
	}
	//告诉模板自定义函数
	t.Funcs(template.FuncMap{
		"autochat":autochat,
	})
	//解析模板
	_, err := t.ParseFiles("./helloV3.tmpl")
	if err != nil {
		fmt.Printf("GoTemplate file parse failed,errMsg:%v\n", err)
		return
	}
	name := "yjl"
	//渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("GoTemplate file render failed,errMsg:%v\n", err)
		return
	}
}
