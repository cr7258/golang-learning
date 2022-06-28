package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var counter Counter
	// 10 个线程读
	for i :=0; i < 10; i++ {
		go func() {
			for {
				fmt.Println(counter.Count()) // 读操作
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 1 个线程写
	for {
		counter.Incr() // 写操作
		fmt.Println("count 计数+1")
		time.Sleep(time.Second)
	}
}

// 一个线程安全的计数器
type Counter struct {
	mu sync.RWMutex
	count int
}

// 使用写锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 使用读锁保护
func (c *Counter) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}