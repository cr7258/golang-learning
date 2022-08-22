package main

import (
	"fmt"
)

/**
 * @description 循环
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

func main() {
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}

	// break 指定层次的循环
	for a := 0; a < 5; a++ {
	outer:
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
				if c == 4 {
					break outer
				}
			}
		}
	}
}
