package main

import "fmt"

/**
结构嵌入
实现自动转发的方法
 */

// Go 语言可以通过结构嵌入实现自动转发方法，在不给定字段名的情况下指定类型即可
type report struct {
	sol int
	temperature
	location
}

type temperature struct {
	high, low celesius
}

type celesius float64

type location struct {
	lat, long float64
}

func (t temperature) average() celesius {
	return (t.high + t.low) / 2
}

func main() {
	report := report{
		sol: 15,
		location: location{-4.5895, 137.4417},
		temperature: temperature{-1.0, -78.0},
	}

	// 可以直接通过 report.average() 来访问 report.temperature.average() 方法
	fmt.Printf("average %v°C\n", report.average())
}
