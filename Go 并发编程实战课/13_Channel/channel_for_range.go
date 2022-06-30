package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// 如果不先关闭 channel，range 读取完数据以后就会卡在那造成 deadlock
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
