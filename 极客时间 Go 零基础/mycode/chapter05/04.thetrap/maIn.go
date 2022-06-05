package main

import "fmt"

/**
 * @description
 * @author chengzw
 * @since 2022/5/26
 */

func main() {
	security := Assets{
		assets: []Asset{
			&GlassDoor{},
			&WoodDoor{},
		},
	}
	fmt.Println("开始上班")
	security.DoStartWork()
	fmt.Println("8小时候，准备下班")
	security.DoStopWork()
	fmt.Println("DONE")
}
