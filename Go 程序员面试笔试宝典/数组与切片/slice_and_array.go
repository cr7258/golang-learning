package main

import "fmt"

/**
数组就是一片连续的内存， slice 实际上是一个结构体，包含三个字段：长度、容量、底层数组。
// runtime/slice.go
type slice struct {
	array unsafe.Pointer // 元素指针
	len   int // 长度
	cap   int // 容量
}

参考题解：https://golang.design/go-questions/slice/vs-array/
*/
func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]     // s1 底层数组 [2,3,4,空,空,空,空,空]
	s2 := s1[2:6:7]      // s2 底层数组 [4,5,6,7,空]，cap(s2)=5；s1[2:6] 直接取到数组末尾， cap(2) = 6
	s2 = append(s2, 100) // s2 底层数组 -> [4,5,6,7,100]  s1 底层数组 -> [2,3,4,空,空,空,100,空]
	s2 = append(s2, 200) // s2 底层数组 -> [4,5,6,7,100,200]（slice 指向新数组）  s1 底层数组 -> [2,3,4,空,空,空,100,空] (不影响)
	s1[2] = 20

	fmt.Println(s1)    // [2,3,20]
	fmt.Println(s2)    // [4,5,6,7,100,200]
	fmt.Println(slice) // [0,1,2,3,20,5,6,7,100,9]
}
