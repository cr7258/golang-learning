package main

import "fmt"

/**
这个函数的参数依次是 元素的类型，老的 slice，新 slice 最小求的容量。
例子中 s 原来只有 2 个元素，len 和 cap 都为 2，append 了三个元素后，长度变为 5，容量最小要变成 5，
即调用 growslice 函数时，传入的第三个参数应该为 5。即 cap=5。
而一方面，doublecap 是原 slice容量的 2 倍，等于 4。满足第一个 if 条件，所以 newcap 变成了 5。

func growslice(et *_type, old slice, cap int) slice {
    // ……
    newcap := old.cap  // 2
	doublecap := newcap + newcap  // 4 = 2 + 2
	if cap > doublecap {  // 5 > 4
		newcap = cap   // 5,  计算出了新容量之后，还没有完，出于内存的高效利用考虑，还要进行内存对齐
	} else {
		// ……
	}
	// ……

参考题解：https://golang.design/go-questions/slice/grow/

*/
func main() {
	s := []int{1, 2}
	s = append(s, 4, 5, 6) // 注意是一次添加 append 3 个元素
	// 结果输出 len=5, cap=6
	// 而不是 len=5, cap=8，如果是分别 append 4, 5, 6，则 cap 是 8
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
}
