package main

import "fmt"

/**
针对常量和字面量的计算将在编译时而不是程序运行时执行，无类型的数值常量将由 big 包提供支持，
所以程序能直接对超过 18 艾的数值常量执行所有的常规运算
*/
func main() {
	const distance = 24000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400

	const days = distance / lightSpeed / secondsPerDay
	//fmt.Println(distance) // 直接打印会引发溢出错误
	fmt.Println("Andromeda Galaxy is ", days, "light days away.")
}
