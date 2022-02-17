package main

import (
	"fmt"
	"math/big"
)

/**
最大无符号整数类型 uint64 存储的数值上限是 18 艾(10^18), 超过这个大小的数可以使用 big 包来处理
*/
func main() {
	// 两种写法等价
	secondsPerDay := big.NewInt(86400)
	fmt.Println(secondsPerDay)

	newSecondsPerDay := new(big.Int)
	newSecondsPerDay.SetString("86400", 10)
	fmt.Println(newSecondsPerDay)
}
