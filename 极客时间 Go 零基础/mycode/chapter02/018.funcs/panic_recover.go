package main

import "fmt"

/*
panic 用于处理严重错误，使当前运行函数直接异常退出。
如果异常退出没有被捕获，则会持续向上层递进，直到有捕获的地方，或 main 函数退出。
*/

func panicAndRecover() {
	defer coverPanicUpgrade()
	var nameScore map[string]int = nil
	nameScore["小强"] = 100
}

// 未能捕获异常
// defer 只关注最后一层函数体，并不关心函数里面的内容
func coverPanic() {
	func() {
		// 捕获异常
		if r := recover(); r != nil {
			fmt.Println("出现严重故障...")
		}
	}()
}

// 捕获异常
func coverPanicUpgrade() {
	if r := recover(); r != nil {
		fmt.Println("出现严重故障...")
	}
}

func main() {
	panicAndRecover()
}
