package main

import "fmt"

/**
 * @description 人把大象装进冰箱，实现接口
 * @author chengzw
 * @since 2022/5/26
 */

type manLegend struct {
}

func (*manLegend) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	fmt.Println("用 manLegend 做 OpenTheDoorOfRefrigerator")
	return nil
}

func (*manLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 manLegend 做 PutElephantIntoRefrigerator")
	return nil
}

func (*manLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manLegend 做 CloseTheDoorOfRefrigerator")
	return nil
}
