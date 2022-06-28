package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	count uint64
}

func main() {
	// 封装好的计数器
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10;i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Incr()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count())
}

// 加 1 的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}