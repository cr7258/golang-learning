package main

import (
	"context"
	"fmt"
	"time"
)

type Options struct {
	Interval time.Duration
}

func reqTask(ctx context.Context, name string){
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(time.Second)
		}
	}
}

func main()  {
	// 创建具有超时通知机制的 Context 对象，设置超时时间为 2s
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	go reqTask(ctx, "worker1")
	go reqTask(ctx, "worker2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
