package main

import "fmt"

func main() {
	type celsius float64
	const degrees = 20
	var temperature celsius = degrees
	temperature += 10
	fmt.Println(temperature)

	var warmUp float64 = 10
	// temperature += warmUp // 类型不匹配报错 mismatched types celsius and float64
	// 需要先将 warmUp 转换成 celsius 类型
	temperature += celsius(warmUp)
	fmt.Println(temperature)
}
