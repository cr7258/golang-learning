package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	fmt.Println("parent lock start")
	mu.Lock()
	fmt.Println("parent locked")
	for i := 0; i <= 2; i++ {
		go func(i int) {
			fmt.Printf("sub(%d) lock start\n", i)
			mu.Lock()
			fmt.Printf("sub(%d) locked\n", i)
			time.Sleep(time.Microsecond * 30)
			mu.Unlock()
			fmt.Printf("sub(%d) unlock\n", i)
		}(i)
	}
	time.Sleep(time.Second * 2)
	mu.Unlock()
	fmt.Println("parent unlock")
	time.Sleep(time.Second * 2)
}
