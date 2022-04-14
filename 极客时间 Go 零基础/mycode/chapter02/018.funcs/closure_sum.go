package main

import "fmt"

/*
闭包
*/

var counter int

func calcSum(nums ...int) (sum int) {
	for _, item := range nums {
		sum += item
	}
	counter++
	return
}

func showUsedTimes() {
	fmt.Println("used: ", counter)
}

func genImprovementFunc() func(percentage float64) {
	base := 1000.0
	// 一个函数和其上下文的引用绑定在一起，被称为闭包
	// 闭包可以让你在内层函数中访问到其外层函数的作用域
	return func(percentage float64) {
		base = base * (1 + percentage)
		fmt.Println(base)
	}
}

func main() {
	fmt.Println(calcSum(13, 3, 43, 56, 65, 43, 76, 2, 4))
	fmt.Println(calcSum(13, 3, 43, 56, 65, 43, 76, 2, 4))
	fmt.Println(calcSum(13, 3, 43, 56, 65, 43, 76, 2, 4))
	showUsedTimes()

	imp := genImprovementFunc() // 返回的是函数
	imp(0.1)
	imp(0.1)
	imp(0.1)
	imp(0.1)
}
