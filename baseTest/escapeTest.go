package main

// 逃逸分析
func main() {
	person := Person{
		Name: "张三",
	}
	person.age = 18
}

type Person struct {
	Name string
	age  int
}
