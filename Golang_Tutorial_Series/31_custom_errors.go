package main

import (
	"fmt"
	"math"
)

/**
 * @description 自定义错误
 * @author chengzw
 * @since 2022/8/17
 * @link
 */

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		//return 0, errors.New("Area calculation failed, radius is less than zero")
		// 使用 Errorf 函数更方便
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}
