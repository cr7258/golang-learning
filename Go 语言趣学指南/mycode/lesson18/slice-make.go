package main

import "fmt"

func main() {
	// 使用 make 函数实行预分配策略可以为底层数组设置初始容量，从而避免额外的内存分配和数据复制操作
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Mercury", "Venus", "Earth")
	fmt.Println(dwarfs)
	fmt.Println(len(dwarfs)) //3
	fmt.Println(cap(dwarfs)) //10

	dwarfs2 := make([]string, 10)
	dwarfs2 = append(dwarfs2, "Mercury", "Venus", "Earth") // 从第 11 个元素开始追加
	fmt.Println(dwarfs2)
	fmt.Println(len(dwarfs2)) //13
	fmt.Println(cap(dwarfs2)) //20
}
