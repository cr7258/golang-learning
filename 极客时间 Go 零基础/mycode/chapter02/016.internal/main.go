package main

import (
	"fmt"
	"reflect"
	"time"
)

// 内置函数
func main() {
	// 获取函数运行时间， defer 在函数运行结束时执行
	// defer 可以运行多个，按照从后往前的顺序执行
	// 即使程序有 panic，defer 也会被调用
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println("运行时间: ", finishTime.Sub(startTime))
	}()
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 2")
	}()

	// 创建一个长度为 3，容量为 4 的数组
	arr2 := make([]string, 3, 4)
	fmt.Printf("len: %d, cap: %d\n", len(arr2), cap(arr2))
	fmt.Println("value:", arr2[0])
	fmt.Println("value:", arr2[1])
	fmt.Println("value:", arr2[2])

	// 返回的是 int 类型对象的指针
	i := new(int)
	fmt.Println(reflect.TypeOf(i))

	// 复制切片，深拷贝
	arr3 := []string{"a", "b", "c"}
	arr4 := make([]string, 3, 6)
	copy(arr4, arr3)
	arr4[2] = "q"
	fmt.Println(arr3)
	fmt.Println(arr4)
}
