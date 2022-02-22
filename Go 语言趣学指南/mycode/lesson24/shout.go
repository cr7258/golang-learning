package main

import (
	"fmt"
	"strings"
)

// martian 和 laser 不需要显式地声明它们实现了一个接口，这是 Go 和 Java 不一样的地方
type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

// shouter 函数能够处理任何一个满足 talker 接口的值，无论它的类型是 martian 还是 laser
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{})
	shout(laser(2))
}
