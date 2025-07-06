package main

import (
	"fmt"
)

func Search() {
	demo := new(Demo)
	DB.Where("name = ?", "test").First(demo)
	fmt.Println(demo.Name.String)
	demos := make([]Demo, 0)
	DB.Where("name = ?", "").Find(&demos)
	fmt.Println(demos)
}
