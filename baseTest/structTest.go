package main

import "fmt"

// struct实现接口
func StructTest() {
	var c Car
	VehileDescribe(&c)
	var b Bike
	VehileDescribe(&b)
}

func VehileDescribe(v Vehile) {
	fmt.Println("交通工具为:", v.Name()+",速度为:", v.Speed(), ",价格为:", v.Price())
}

type Vehile interface {
	Speed() string
	Price() int64
	Name() string
}

var _ Vehile = (*Car)(nil)

type Car struct {
}

func (c *Car) Speed() string {
	return "100km/h"
}

func (c *Car) Price() int64 {
	return 100000
}

func (c *Car) Name() string {
	return "小轿车"
}

type Bike struct {
}

func (c *Bike) Speed() string {
	return "30km/h"
}

func (c *Bike) Price() int64 {
	return 1000
}

func (c *Bike) Name() string {
	return "自行车"
}
