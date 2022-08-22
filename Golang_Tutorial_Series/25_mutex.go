package main

import (
	"fmt"
	"sync"
)

/**
 * @description 互斥锁
 * @author chengzw
 * @since 2022/8/16
 * @link
 */
var x = 0

// 要传递指针
// 如果互斥对象是通过值传递的，而不是通过地址传递的，那么每个 Goroutine 都会有自己的互斥对象副本，竞态条件仍然会发生。
func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
