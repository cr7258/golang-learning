package main

import (
	"fmt"
	"sync"
)

// Visited 用于记录网页是否被访问过
// 它的方法可以在多个 goroutine 中并发使用
type Visited struct {
	// mu 用于保护 visited 映射
	mu      sync.Mutex     // 声明一个互斥锁
	visited map[string]int // 声明一个从网址（字符串）键指向整数值的映射
}

// VisitLink 会记录本次针对给定网址的访问，然后返回更新之后的链接统计值
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()         // 锁定互斥锁
	defer v.mu.Unlock() // 确保锁会在方法执行完毕之后解除
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
	visited := &Visited{
		mu:      sync.Mutex{},
		visited: map[string]int{},
	}
	for i := 0; i < 100; i++ {
		go visited.VisitLink("www.baidu.com")
		//time.Sleep(time.Microsecond)
	}
	fmt.Println(visited.VisitLink("www.baidu.com"))
}
