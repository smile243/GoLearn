package main

import "gorm.io/gorm"

func Delete(db *gorm.DB) {
	//如果没有主键，会逻辑删除所有数据
	var demo = Demo{}
	demo.Age = 18
	db.Delete(&demo)
	//db.Where("name = ?", "test").Delete(&Demo{})
}
