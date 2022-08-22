package main

/**
 * @description 死锁
 * @author chengzw
 * @since 2022/8/16
 * @link
 */

func main() {
	ch := make(chan int)
	ch <- 5 // 没有 channel 在等待读取数据，会导致 panic
}
