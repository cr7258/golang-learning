package _2_prime

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 启动两个 goroutine
func TestRunPrime3(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("第一个worker开始计算", time.Now())
		result = append(result, collectPrime(2, 100000)...)
		fmt.Println("第一个worker完成计算", time.Now())
	}()
	go func() {
		defer wg.Done()
		fmt.Println("第二个worker开始计算", time.Now())
		result = append(result, collectPrime(100001, 200000)...)
		fmt.Println("第二个worker完成计算", time.Now())
	}()
	wg.Wait()
	finishTime := time.Now()
	fmt.Println("finishTime: ", finishTime)
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}

func collectPrime(start int, end int) (result []int) {
	for num := start; num <= end; num++ {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	return
}

func isPrime(num int) (isPrime bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			return
		}
	}
	isPrime = true
	return
}
