package main

import (
	"fmt"

	"gorm.io/gorm"
)

func Search(db *gorm.DB) {
	demo := new(Demo)
	db.Where("name = ?", "test").First(demo)
	fmt.Println(demo.Name.String)
	demos := make([]Demo, 0)
	db.Where("name = ?", "").Find(&demos)
	fmt.Println(demos)
}
