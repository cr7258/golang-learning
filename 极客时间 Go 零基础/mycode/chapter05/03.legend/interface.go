package main

/**
 * @description 接口
 * @author chengzw
 * @since 2022/5/26
 */
type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant, Refrigerator) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
}

type Refrigerator struct {
	Size string
}

type Elephant struct {
	Name string
}
