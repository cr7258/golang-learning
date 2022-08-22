package main

import "fmt"

/**
 * @description 方法中的值接收者与函数中的值参数
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p)
	//area(*p) 可以
	p.area() //calling value receiver with a pointer
}
