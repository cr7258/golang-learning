package main

import "fmt"

func main() {
	var administrator *string
	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)
	bolden = "Charles Frank Bolden Jr." // 修改 bolden 的值会影响的 *administrator 引用值
	fmt.Println(*administrator)
	*administrator = "Maj. Gen. Charles Frank Bolden Jr." // 可以通过修改解引用 *administrator 来间接修改 bolden 的值
	fmt.Println(bolden)
	fmt.Println(administrator == &bolden) // 指向的内存地址一致

	// 解引用的值赋值给另一个变量将产生一个字符串副本，修改将不会互相影响
	charles := *administrator
	*administrator = "Charless Bolden"
	fmt.Println(charles)
	fmt.Println(&charles == administrator) // 内存地址不一致
}
