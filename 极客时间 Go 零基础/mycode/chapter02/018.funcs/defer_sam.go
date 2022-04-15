package main

import (
	"fmt"
	"os"
	"time"
)

/*
defer 在函数运行结束时执行
*/
func openFile() {
	fileName := "/text.txt"
	fileOjb, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	defer fileOjb.Close()
}

func deferGuess() {
	startTime := time.Now()
	defer fmt.Println("duration: ", time.Now().Sub(startTime)) // 等了 5s 才输出，但是最后的 duration 是 几 ns，defer 虽然是最后才运行的，但是函数已经预先准备好了
	// 如果想要正确计算延时，可以使用闭包
	defer func() {
		fmt.Println("duration closure: ", time.Now().Sub(startTime))
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("start time: ", startTime)
}

func main() {
	// openFile()
	deferGuess()
}
