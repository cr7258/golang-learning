package main

import "fmt"

/**
 * @description 切片
 * @author chengzw
 * @since 2022/8/15
 * @link
 */
func main() {
	// 根据原有 array 或者 slice 创建新的 slice
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)

	// 直接声明 slice，不填写元素个数
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)

	// 使用 make 创建切片
	// make([]T, len, cap)
	i := make([]int, 5, 5)
	fmt.Println(i)

	// 往切片中添加元素
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	// 往切片中添加切片，使用 ...
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	// 复制切片
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) // 复制完以后可以使用新的切片，GC 就可以对原始数组进行回收
	fmt.Println(countriesCpy)
}
