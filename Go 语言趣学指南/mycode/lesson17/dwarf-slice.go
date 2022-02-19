package main

import "fmt"

func main() {
	// 直接声明切片，使用 []string 作为类型
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// 什么数组
	dwarfsArray := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Printf("%T\n", dwarfs)
	fmt.Printf("%T\n", dwarfsArray)
}
