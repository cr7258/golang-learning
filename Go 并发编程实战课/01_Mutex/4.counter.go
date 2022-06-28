package main

import (
	"fmt"
	"sync"
)

/**
采用嵌入字段的方式。通过嵌入字段，你可以在这个 struct 上直接调用 Lock/Unlock 方法。
 */

type Count struct {
	 sync.Mutex
	 Count uint64
}

func main(){
	var count Count
	var wg sync.WaitGroup
	wg.Add(10)

	for i :=0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j :=0; j < 100000; j++{
				count.Lock()
				count.Count++
				count.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(count.Count)
}