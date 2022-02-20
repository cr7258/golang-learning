package main

import "fmt"
/**
合并结构
 */

type temperature struct {
	high, low celesius
}

type celesius float64

type location struct {
	lat, long float64
}

type report struct {
	sol         int
	temperature temperature
	location    location
}

func (t temperature) average() celesius {
	return (t.high + t.low) / 2
}

// 转发方法
func (r report) average() celesius{
	return r.temperature.average()
}
func main(){
	t := temperature{high: -1.0, low: -78.0}
	fmt.Printf("average %v°C\n", t.average())

	report := report{sol: 15, temperature: t}
	fmt.Printf("average %v°C\n", report.temperature.average())
	fmt.Printf("average %v°C\n", report.average()) // 转发方法

}
