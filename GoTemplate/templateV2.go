package GoTemplate

import (
	"fmt"
	"html/template"
	"net/http"
)

func SayHelloV2(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("GoTemplate/helloV2.tmpl")
	if err != nil {
		fmt.Printf("GoTemplate file parse failed,errMsg:%v\n", err)
		return
	}
	user := UserInfo{
		Name:   " zs ",
		Gender: "男",
		Age:    18,
	}
	//interface{}代表任意类型
	m1 := map[string]any{
		"Name":   "ls",
		"Gender": "男",
		"Age":    19,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"乒乓球",
	}
	//渲染模板
	err = t.Execute(w, map[string]any{
		"u1":     user,
		"m1":     m1,
		"hobbys": hobbyList,
	})
	if err != nil {
		fmt.Printf("GoTemplate file render failed,errMsg:%v\n", err)
		return
	}
}
