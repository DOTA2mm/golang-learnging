package main

import (
	"fmt"
	"time"
)

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		// for 循环的 range 语句可以用在通道 ch 上
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func main() {
	suck(pump())
	time.Sleep(1e9)
}
