package main

import "fmt"

func main() {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	//terrestrial := planets[0:4:4] // 长度 4，容量 4 的切片
	terrestrial := planets[0:4] // 长度 4，容量 8 的切片
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)
}
