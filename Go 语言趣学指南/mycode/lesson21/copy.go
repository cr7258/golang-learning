package main

import "fmt"

func main() {
	// 数组复制是一个完整的副本，修改新数组不会影响原先的数组
	array1 := [3]string{"apple", "banana", "orange"}
	array2 := array1
	array2[1] = "watermelon"
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println("=========================")
	// 修改切片会影响原先的数组
	slice1 := [3]string{"apple", "banana", "orange"}
	slice2 := slice1[:]
	slice2[1] = "watermelon"
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("=========================")
	// 修改映射会影响原先的映射
	map1 := map[string]string{
		"banana": "yellow",
		"apple":  "red",
	}
	map2 := map1
	map2["apple"] = "newRed"
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println("=========================")
	// 修改结构不会影响原先的结构
	type fruit struct {
		banana, apple string
	}
	struct1 := fruit{banana: "yellow", apple: "red"}
	struct2 := struct1
	struct2.apple = "newRed"
	fmt.Println(struct1)
	fmt.Println(struct2)

}
