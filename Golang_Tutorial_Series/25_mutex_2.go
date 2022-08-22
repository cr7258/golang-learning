package main

import (
	"fmt"
	"sync"
)

/**
 * @description 使用 channel 解决竞态条件
 * @author chengzw
 * @since 2022/8/16
 * @link
 */
var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1) // 创建一个容量为 1 的 channel，确保只有一个 Goroutine 能够访问
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
