package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew", int(*l))
}

func main() {
	// 无论是 martian 还是指向 martian 的指针，都可以满足 talker 接口
	// 如果类型的非指针版本能够满足接口，那么它的指针版本也能够满足
	shout(martian{})
	shout(&martian{})

	pew := laser(2)
	shout(&pew) // 传入 pew 无法满足接口
}
