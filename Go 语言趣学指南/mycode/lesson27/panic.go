package main

import "fmt"

func main() {
	var nowhere *int
	fmt.Println(nowhere) // nil
	//fmt.Println(*nowhere) // 空指针异常
}
