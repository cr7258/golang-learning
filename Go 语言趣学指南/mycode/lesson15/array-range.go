package main

import "fmt"

func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for i, draft := range dwarfs {
		fmt.Println(i, draft)
	}
}
