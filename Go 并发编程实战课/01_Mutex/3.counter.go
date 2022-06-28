package main

import (
	"fmt"
	"sync"
)

/**
将 Mutex 嵌入到 struct 中使用
 */

type Count struct {
	 mu sync.Mutex
	 Count uint64
}

func main(){
	var count Count
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j :=0; j < 100000; j++ {
				count.mu.Lock()
				count.Count++
				count.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count.Count)
}