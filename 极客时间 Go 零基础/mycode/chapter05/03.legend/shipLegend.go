package main

import "fmt"

/**
 * @description 用船装大象，实现接口
 * @author chengzw
 * @since 2022/5/26
 */
type shipLegend struct {
}

func (*shipLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 ship 做 OpenTheDoorOfRefrigerator")
	return nil
}
func (*shipLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 ship 做 PutElephantIntoRefrigerator")
	return nil
}
func (*shipLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 ship 做 CloseTheDoorOfRefrigerator")
	return nil
}
