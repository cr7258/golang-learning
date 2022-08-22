package main

import "fmt"

/**
 * @description 函数
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

// 最简单的函数
func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// 多返回值
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// 命名返回值
func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

	area2, perimeter2 := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area2, perimeter2)
}
