package main

import "fmt"

/*
作用域
*/

// 全局变量
var tall float64
var weight float64

func main() {
	fmt.Println("全局变量赋值前: ", calcAdd()) // 0 + 0 = 0

	tall, weight = 1.70, 70.0
	fmt.Println("全局变量赋值后: ", calcAdd()) // 71.7

	tall, weight := 100.00, 70.00
	fmt.Println(tall, weight)               // 100.00, 70.00
	fmt.Println("重新定义重名的局部变量: ", calcAdd()) // 71.7, 不会影响全局变量

	calculatorAdd := func(a, b int) int {
		return a + b
	}

	// Golang 中使用 {} 来定义作用域
	{
		personTall := 181
		personWeight := 90
		calculatorAdd(int(personTall), personWeight)
	}

	result := calculatorAdd(1, 3)
	fmt.Println(result)

	sampleSubdomain()
	sampleSubdomainIf()
}

func sampleSubdomain() {
	name := "小强"
	fmt.Println("名字是: ", name) // 小强
	{
		fmt.Println("名字是: ", name) // 小强
		name = "小王"                // 重新赋值，覆盖母作用域的值
		//name := "小王"       // 如果是 := 是重新定义了一个新的局部变量 name， 并且变量的有效范围是 {} 定义的作用域内，这样不会影响作用域外变量的值
		fmt.Println("名字是: ", name) // 小王
	}
	fmt.Println("名字是: ", name) // 小王
}

func calcAdd() float64 {
	return tall + weight
}
func sampleSubdomainIf() {
	if v := calcAdd(); v == 0 {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
	//fmt.Println(v) // 无效 。v的有效范围为 if block
}
