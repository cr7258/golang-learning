package main

import (
	"fmt"
	"sync"
)

/**
Go 的 Mutex 不支持重入，是不可重入锁。
因为 Mutex 的实现中没有记录哪个 goroutine 拥有这把锁。理论上，任何 goroutine 都可以随意地 Unlock 这把锁，所以没办法计算重入条件，
 */


func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}


func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}


func main() {
	l := &sync.Mutex{}
	foo(l)
}