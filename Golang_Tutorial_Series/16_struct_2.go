package main

import "fmt"

/**
 * @description 匿名结构体
 * @author chengzw
 * @since 2022/8/15
 * @link
 */
type Address struct {
	city  string
	state string
}
type Person struct {
	name string
	age  int
	Address
}

func main() {
	p := Person{
		name: "Naveen",
		age:  50,
		Address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field
}
