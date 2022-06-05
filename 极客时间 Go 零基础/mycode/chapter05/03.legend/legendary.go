package main

import "fmt"

/**
 * @description
 * @author chengzw
 * @since 2022/5/26
 */

func legendary(legend PutElephantIntoRefrigerator, r Refrigerator, e Elephant) {
	fmt.Println("传说中，大象可以这么装...")

	// human 警告
	if _, ok := legend.(*manLegend); ok {
		fmt.Println("WARNING：现在还在用人工，效率太低")
	}
	legend.OpenTheDoorOfRefrigerator(r)
	legend.PutElephantIntoRefrigerator(e, r)
	legend.CloseTheDoorOfRefrigerator(r)

	fmt.Println("this is a legendary")
}
