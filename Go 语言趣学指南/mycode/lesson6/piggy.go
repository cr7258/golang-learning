package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var price float64
	for price < 20.00 {
		switch rand.Intn(3) {
		case 0:
			price += 0.05
		case 1:
			price += 0.1
		case 2:
			price += 0.25
		}
		fmt.Printf("当前余额: %5.2f\n", price)
	}
}
