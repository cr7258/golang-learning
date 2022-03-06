package main

import (
	"fmt"
	"time"
)

func sleepyGophere(i int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("%v... snore ...\n", i)
}

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGophere(i)
	}
	time.Sleep(4 * time.Second) // 所有 goroutine 将在主线程运行完毕后终止
}
