package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// (err error)是给函数的返回命名，这样可以直接在函数体内使用变量
func initMySql() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//修改默认表名复数形式
			//SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	initMySql()
	DB.AutoMigrate(&Demo{})
	//CreateZeroValue()
	//Search()
	//Update()
	//Delete()
}
