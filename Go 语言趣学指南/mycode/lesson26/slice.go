package main

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	planetes := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto",
	}

	// reclassify 函数修改切片长度，传入的是指针
	reclassify(&planetes)
	fmt.Println(planetes)
}
