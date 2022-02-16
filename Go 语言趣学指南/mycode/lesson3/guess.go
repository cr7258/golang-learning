package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = 42
	for {
		var n = rand.Intn(100) + 1
		if n < num {
			fmt.Printf("%v is too small\n", n)
		} else if n > num {
			fmt.Printf("%v is too big\n", n)
		} else {
			fmt.Printf("%v you got it!\n", n)
			break
		}
	}
}
