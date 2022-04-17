package main

import (
	"fmt"
	"runtime/debug"
)

func convertType() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic...")
			debug.PrintStack()
		}
	}()
	var a interface{}
	a = "string aaa"

	b := a.(int)
	fmt.Println(b)
}

func main() {
	convertType()
	fmt.Println("finish")
}
