package main

import (
	"fmt"
	"sync"
)

/**
修改 counter 变量时加锁，修改完后释放锁，可以保证正确修改共享变量
*/

func main() {
	// 互斥锁保护计数器
	// Mutex 的零值是还没有 goroutine 等待的未加锁的状态，所以你不需要额外的初始化，直接声明变量（如 var mu sync.Mutex）即可。
	var mu sync.Mutex
	// 计数器的值
	var count = 0

	// 辅助变量，用来确认所有的 goroutine 都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动 10 个 goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加 10w 次
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
