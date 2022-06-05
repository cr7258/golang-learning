package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//countDict()
	//countDictGoRoutineSafe()
	countDictGoroutineSafeRW()
	//countDictForgetUnlock()
}

// 有重复覆盖的问题
func countDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			totalCount += 100 // 注意，这里有重复覆盖的问题
		}()
	}
	wg.Wait()
	fmt.Println("预计有: ", 100*5000, " 字")
	fmt.Println("总共有: ", totalCount, " 字")
}

// 加锁，不会重复覆盖，正确写法
func countDictGoRoutineSafe() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	// 锁
	totalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			totalCountLock.Lock()
			defer totalCountLock.Unlock()
			totalCount += 100

		}()
	}
	wg.Wait()
	fmt.Println("预计有: ", 100*5000, " 字")
	fmt.Println("总共有: ", totalCount, " 字")
}

// 读写锁
func countDictGoroutineSafeRW() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.RWMutex{}

	wg := sync.WaitGroup{}
	workerCount := 5
	wg.Add(workerCount)

	doneCh := make(chan struct{})
	for p := 0; p < workerCount; p++ {
		go func(p int) { // 读锁可以多个go routine同时拿到。
			fmt.Println(p, "读锁开始时间：", time.Now())
			totalCountLock.RLock()
			fmt.Println(p, "读锁拿到锁时间：", time.Now())
			time.Sleep(1 * time.Second)
			totalCountLock.RUnlock()
		}(p)
	}
	for p := 0; p < workerCount; p++ {
		go func(p int) {
			defer wg.Done()
			fmt.Println(p, "写锁开始时间：", time.Now())
			totalCountLock.Lock()
			fmt.Println(p, "写锁拿到锁时间：", time.Now())
			defer totalCountLock.Unlock()
			totalCount += 100
		}(p)
	}
	wg.Wait()
	close(doneCh)
	time.Sleep(1 * time.Second)
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}

// 忘记释放锁，导致 deadlock
func countDictForgetUnlock() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go func() {
			defer wg.Done()
			totalCountLock.Lock()
			totalCount += 100
			// 忘记释放锁
			// totalCountLock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}
