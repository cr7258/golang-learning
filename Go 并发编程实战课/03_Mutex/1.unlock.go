package main

import (
	"fmt"
	"sync"
)

/**
错误案例 1：Lock/Unlock 不是成对出现
 */
func main() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world!")
}
