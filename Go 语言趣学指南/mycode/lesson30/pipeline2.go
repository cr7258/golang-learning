package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	// 使用 close 关闭通道
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	//for {
	//	item, ok := <-upstream
	//	if !ok {
	//		close(downstream)
	//		return
	//	}
	//	if !strings.Contains(item, "bad") {
	//		downstream <- item
	//	}
	//}
	// 通过 range 循环以更简单的方式处理通道数据，效果和上面一致
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
