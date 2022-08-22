package main

import "fmt"

/**
* @description
  Go 里面 switch 默认相当于每个 case 最后带有 break，匹配成功后不会自动向下执行其他 case，
  而是跳出整个 switch, 但是可以使用 fallthrough 强制执行后面的 case 代码，不管后面的结果是否满足 case
* @author chengzw
* @since 2022/8/15
* @link
*/
func number() int {
	num := 15 * 5
	return num
}

func main() {

	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 20: // 不满足也会执行
		fmt.Printf("%d is lesser than 20", num)
	}
}
