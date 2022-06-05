package main

import "fmt"

/**
 * @description
 * @author chengzw
 * @since 2022/5/26
 */

// 确保完整实现接口
var _ Door = &GlassDoor{}

type GlassDoor struct {
}

func (d *GlassDoor) Unlock() {
	fmt.Println("GlassDoor Unlock")
}

func (d *GlassDoor) Lock() {
	fmt.Println("GlassDoor Lock")
}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}
func (*GlassDoor) Close() {
	fmt.Println("GlassDoor Close")
}
