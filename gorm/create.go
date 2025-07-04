package main

import (
	"database/sql"

	"gorm.io/gorm"
)

func Create(db *gorm.DB) {
	//name := "test"
	name := sql.NullString{String: "", Valid: true}
	db.Create(&Demo{Name: name, Age: 18})
}

func CreateZeroValue(db *gorm.DB) {
	//方式一
	//db.Create(&Demo{Name: new(string), Age: 18})
	//方式二
	db.Create(&Demo{Name: sql.NullString{String: "", Valid: true}, Age: 18})
}
