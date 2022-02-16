package main

import "fmt"

func main() {
	const hoursPerDay = 24
	var days = 28
	var distance = 56000000

	fmt.Println(distance/days*hoursPerDay, "km/h")
}
