package main

import (
	"gorm.io/gorm"
)

func Update(db *gorm.DB) {
	//save默认修改所有字段
	//更新多个记录
	db.Model(&Demo{}).Where("name = ?", "test").UpdateColumn("name", "test2")
}
