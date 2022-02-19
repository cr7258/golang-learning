package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor 函数类型
type sensor func() kelvin

func realSensor() kelvin {
	return 1
}

func calibrate(s sensor, offset kelvin) sensor {
	// 声明并返回匿名函数
	return func() kelvin {
		return s() + offset
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	var num kelvin = 7
	sensor := calibrate(realSensor, num)
	fmt.Println(sensor()) //7+1=8
	num++                 //修改 num 并不会修改 t() 的结果，因为 offset 接受的是实参的副本值而不是引用，也就是俗称的传值
	fmt.Println(sensor()) //8

	// 多次调用会每次都会调用 fakeSensor 函数并生成新的随机数
	fakeSensor := calibrate(fakeSensor, 5)
	fmt.Println(fakeSensor())
	fmt.Println(fakeSensor())
	fmt.Println(fakeSensor())
}
