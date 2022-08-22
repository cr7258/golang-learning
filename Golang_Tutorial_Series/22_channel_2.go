package main

import "fmt"

/**
 * @description
 * @author chengzw
 * @since 2022/8/16
 * @link
 */
func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}
func main() {
	number := 589
	squareop := make(chan int)
	cubeop := make(chan int)
	go calcSquares(number, squareop)
	go calcCubes(number, cubeop)
	squares, cubes := <-squareop, <-cubeop
	fmt.Println("Final output", squares+cubes)
}
