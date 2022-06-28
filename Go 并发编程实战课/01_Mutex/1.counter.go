package main

import (
	"fmt"
	"sync"
)

/**
10 个 goroutine 同时对 count 变量操作，没有加锁，会产生并发问题，最后结果小于 1000000
*/
func main() {
	var count = 0
	// 使用 WaitGroup 等待 10 个 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动 10 个 goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加 10w 次
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}

	// 等待 10 个 goroutine 完成
	wg.Wait()
	fmt.Println(count)
}
