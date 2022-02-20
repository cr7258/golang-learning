package main

import (
	"fmt"
	"strings"
)

/**
任何类型的任何值，只要它满足了接口的要求，就能成为变量 t 的值。
 */

// 定义一个接口
var t interface{
	talk() string
}

// martian 和 laser 不需要显式地声明它们实现了一个接口，这是 Go 和 Java 不一样的地方
type martian struct {}

func (m martian) talk() string{
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main()  {
	// 多态
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
}
