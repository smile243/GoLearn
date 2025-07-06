package main

import (
	"database/sql"
)

func Create() {
	//name := "test"
	name := sql.NullString{String: "", Valid: true}
	DB.Create(&Demo{Name: name, Age: 18})
}

func CreateZeroValue() {
	//方式一
	//db.Create(&Demo{Name: new(string), Age: 18})
	//方式二
	DB.Create(&Demo{Name: sql.NullString{String: "", Valid: true}, Age: 18})
}
