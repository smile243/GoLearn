package main

import (
	"database/sql"

	"gorm.io/gorm"
)

type Demo struct {
	ID int32
	//方式一
	//Name *string `gorm:"type:varchar(50);default:yjl"`
	//方式二
	Name sql.NullString `gorm:"type:varchar(50);default:yjl"`
	Age  int32
}

var (
	DB *gorm.DB
)
