package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @description 优雅关闭 channle：N 个 sender，一个 reciver
 * @author chengzw
 * @since 2022/7/1
 * @link https://golang.design/go-questions/channel/graceful-close/
 */

func main() {
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// receiver
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)
				return
			}
			fmt.Println(value)
		}
	}()

	select {
	case <-time.After(time.Hour):
	}
}
