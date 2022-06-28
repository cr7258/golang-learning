package main

/**
Go 语言原生 map 并不是线程安全的，对它进行并发读写操作的时候，需要加锁。
而 sync.map 则是一种并发安全的 map，在 Go 1.9 引入。

https://www.cnblogs.com/qcrao-2018/p/12833787.html
 */


import (
	"fmt"
	"sync"
)

func main()  {
	var m sync.Map
	// 1. 写入
	m.Store("qcrao", 18)
	m.Store("stefno", 20)

	// 2. 读取
	age, _ := m.Load("qcrao")
	fmt.Println(age.(int))

	// 3. 遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})

	// 4. 删除
	m.Delete("qcrao")
	age, ok := m.Load("qcrao")
	fmt.Println(age, ok)

	// 5. 读取或写入
	m.LoadOrStore("stefno", 100)
	age, _ = m.Load("stefno")
	fmt.Println(age)
}