package main

// goroutine是通过Go的runtime管理的一个线程管理器
// goroutine通过go关键字实现了，其实就是一个普通的函数

import "fmt"

// 定义一个channel时，也需要定义发送到channel的值的类型
// 注意，必须使用make 创建channel

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}
func main() {
	a := []int{7, 2, 8, 9, -3, 1}
	c := make(chan int)     // channel 必须由 make 创建
	go sum(a[:len(a)/2], c) // 截取后半部分
	go sum(a[len(a)/2:], c)

	// 如果读取（value := <-ch）它将会被阻塞，直到有数据接收
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
	// ch := make(chan type, value)
	// 当 value = 0 时，channel 是无缓冲阻塞读写的，
	// 当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入
	ch := make(chan int, 1)
	ch <- 1 // 发送
	ch <- 2
	fmt.Println(<-ch) // 接收
	fmt.Println(<-ch)
}
