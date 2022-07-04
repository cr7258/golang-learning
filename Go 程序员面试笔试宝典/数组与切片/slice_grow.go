package main

/**
slice 扩容策略：
1.18 版本之前：当原 slice 容量小于 1024 的时候，新 slice 容量变成原来的 2 倍；原 slice 容量超过 1024，新 slice 容量变成原来的 1.25 倍。
1.18 版本之后：当原 slice 容量 (oldcap) 小于 256 的时候，新 slice(newcap) 容量为原来的 2 倍；原 slice 容量超过 256，新 slice 容量 newcap = oldcap+(oldcap+3*256)/4
*/

import "fmt"

func main() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}
