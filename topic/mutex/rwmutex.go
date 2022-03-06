package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwm sync.RWMutex

	for i := 0; i <= 2; i++ {
		go func(i int) {
			fmt.Printf("go(%d) start lock\n", i)
			rwm.RLock()
			fmt.Printf("go(%d) locked\n", i)
			time.Sleep(time.Second * 2)
			rwm.RUnlock()
			fmt.Printf("go(%d) unlock\n", i)
		}(i)
	}
	// 先 sleep 一小会，保证 for的 goroutine 都会执行
	time.Sleep(time.Microsecond * 100)
	fmt.Println("main start lock")
	// 当子进程都执行时，且子进程所有的资源都已经 Unlock 了
	// 父进程才会执行
	rwm.Lock()
	fmt.Println("main locked")
	time.Sleep(time.Second)
	rwm.Unlock()
}
