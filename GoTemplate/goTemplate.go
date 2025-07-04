package GoTemplate

import (
	"fmt"
	"net/http"
)

func GoTemplate() {
	http.HandleFunc("/index", Index)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/v2", SayHelloV2)
	http.HandleFunc("/v3", sayHelloV3)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http server start failed,errMsg:%v\n", err)
		return
	}
}
