package main

import (
	"fmt"
	"time"
)

func init() {
	// 同步通道-使用带缓冲的通道， 50 为容量
	c := make(chan int, 50)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}
