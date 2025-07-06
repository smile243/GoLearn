package main

func Update() {
	//save默认修改所有字段
	//更新多个记录
	DB.Model(&Demo{}).Where("name = ?", "test").UpdateColumn("name", "test2")
}
