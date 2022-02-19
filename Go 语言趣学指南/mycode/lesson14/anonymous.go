package main

import "fmt"

//在包中已经声明过了
//type kelvin float64

func main() {
	var k kelvin = 294.0
	sensor := func() kelvin {
		return k
	}
	fmt.Println(sensor()) //294
	k++
	fmt.Println(sensor()) //295
}
