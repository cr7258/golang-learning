package main

import (
	"fmt"
)

/*
递归
*/
func guess(left, right int) {
	guessed := (left + right) / 2
	var getFromInput string
	fmt.Println("我猜是：", guessed)
	fmt.Println("如果高了, 输入 1; 如果低了, 输入 0; 对了, 输入 9")
	fmt.Scanln(&getFromInput)
	switch getFromInput {
	case "1":
		if left == right {
			fmt.Println("你是不是改主意了?")
			return
		}
		guess(left, guessed-1)
	case "0":
		if left == right {
			fmt.Println("你是不是改主意了?")
			return
		}
		guess(guessed+1, right)
	case "9":
		fmt.Println("你心里想的数字是: ", guessed)
		return
	}
}

func main() {
	guess(1, 100)
}
