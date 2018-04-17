package main

import "fmt"

// 存在多个channel的时候，我们该如何操作呢，
// Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。

func fibonacci1(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func init() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci1(c, quit)
}
