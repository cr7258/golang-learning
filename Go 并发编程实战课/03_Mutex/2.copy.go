package main

import (
	"fmt"
	"sync"
)

/**
错误案例 2：复制了一个使用了的 Mutex，导致锁无法使用，程序处于死锁的状态。
可以使用 go vet 2.copy.go 命令发现 Mutex 复制的问题
 */

type Counter struct {
	sync.Mutex
	Count int
}


func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}