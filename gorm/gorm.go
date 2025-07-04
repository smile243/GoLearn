package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//修改默认表名复数形式
			//SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Demo{})
	//CreateZeroValue(db)
	//Search(db)
	//Update(db)
	Delete(db)
}
