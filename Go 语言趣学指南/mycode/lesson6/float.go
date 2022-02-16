package main

import (
	"fmt"
	"math"
)

/**
浮点数有有误差
*/
func main() {
	piggyBank := 0.1
	piggyBank += 0.2       // 0.30000000000000004
	fmt.Println(piggyBank) // false

	// 正确比较的方法
	fmt.Println(math.Abs(piggyBank-0.3) < 0.00001) // true
}
