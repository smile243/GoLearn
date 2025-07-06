package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// func main() {
// 	//slice()
// 	//MD5()
// 	//channelWithoutBuffer()
// 	//deferTest()
// 	StructTest()
// }

func slice() {
	sli := make([]int, 2, 5)
	fmt.Println(len(sli), cap(sli))
	//删除尾部 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[:len(sli)-2]), cap(sli[:len(sli)-2]), sli[:len(sli)-2])
}

type User struct {
	Id   int
	Name string
}

// MD5 方法
func MD5() {
	u := User{
		Id:   1,
		Name: "张三",
	}
	s := md5.New()
	// 将整个结构体转换为字符串
	data := fmt.Sprintf("%d-%s", u.Id, u.Name)
	s.Write([]byte(data))
	fmt.Printf("User对象的MD5值(%+v): %s\n", u, hex.EncodeToString(s.Sum(nil)))
}

func channelWithoutBuffer() {
	fmt.Println("main start")
	//不带缓冲区，读写都会阻塞
	ch := make(chan int)
	//会发生死锁，因为ch<-1会阻塞住，轮不到go func执行
	// ch <- 1
	// go func() {
	// 	fmt.Println(<-ch)
	// }()
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
	fmt.Println("main end")
}

func deferTest() {
	//defer会在函数return之前调用
	t5()
}

func t2() (a int) {
	//闭包，引用传递
	defer func() {
		a++
	}()
	return 1
}

func t4() (a int) {
	//值传递
	defer func(a int) {
		a++
	}(a)
	return 1
}

func t5() {
	x := 1
	y := 2
	defer calc("A", x, calc("B", x, y))
	x = 3
	defer calc("C", x, calc("D", x, y))
	y = 4
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
