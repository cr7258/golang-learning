package main

import "fmt"

/**
 * @description
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

func main() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}

	newEmp := "joe"
	// 检查 map 中是否含有指定 key
	if value, ok := employeeSalary[newEmp]; ok {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")

	// 遍历 map
	for key, value := range employeeSalary {
		fmt.Println(key, value)
	}

	// 从 map 中删除元素
	delete(employeeSalary, "steve")
	fmt.Println(employeeSalary)

	// map 是引用类型，当 map 分配给新的变量时，它们都指向相同的内存地址，修改其中一个会影响另一个的值
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary)
}
