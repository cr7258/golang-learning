package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

// 地球日期
// 标准库中的 time.Time 类型满足了 stardater 接口，所以 stardate 函数能够继续处理地球的日期
func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

// 火星日期
// sol 类型通过实现 yearDay() 方法和 Hour() 方法满足了 stardater 接口
type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

func main() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate(s))
}
