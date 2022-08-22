package main

import "fmt"

/**
 * @description 结构体和指针
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)

	// Go 允许我们不显式通过 * 取消引用
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)
}
