package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "A"
	ch <- "B"
	ch <- "C"
	ch <- "D"
	ch <- "E"
	close(ch)
}

func getData(ch chan string) {
	// for {
	// 	// 检测 channel 是否被关闭
	// 	input, open := <-ch
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Printf("%s ", input)
	// }

	// for - range /自动检测通道是否关闭
	for input := range ch {
		fmt.Printf("%s ", input)
	}
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}
