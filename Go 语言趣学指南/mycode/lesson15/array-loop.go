package main

import "fmt"

func main() {
	// [] 中可以使用 ...，而不是具体数字作为数组长度，让 Go 编译器自动计算
	dwarfs := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}
}
