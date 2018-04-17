package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 生产者通过内置函数close关闭channel
	// 在消费方可以通过语法v, ok := <-ch测试channel是否被关闭
	// 如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭
	close(c)
}

func init() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// 通过 rang 操作缓存类型的channel
	for i := range c {
		fmt.Println(i)
	}
}
