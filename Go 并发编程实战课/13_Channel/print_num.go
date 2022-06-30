package main

import (
	"fmt"
	"time"
)

/**
有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，
要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
*/
func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	ch4 := make(chan bool)

	go func() {
		for {
			<-ch1
			fmt.Println("I'm goroutine 1")
			time.Sleep(time.Second)
			ch2 <- true
		}
	}()

	go func() {
		for {
			<-ch2
			fmt.Println("I'm goroutine 2")
			time.Sleep(time.Second)
			ch3 <- true
		}
	}()

	go func() {
		for {
			<-ch3
			fmt.Println("I'm goroutine 3")
			time.Sleep(time.Second)
			ch4 <- true
		}
	}()

	go func() {
		for {
			<-ch4
			fmt.Println("I'm goroutine 4")
			time.Sleep(time.Second)
			ch1 <- true
		}
	}()

	ch1 <- true
	select {}
}
