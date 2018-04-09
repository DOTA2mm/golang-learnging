package main

import "fmt"

// defer是采用后进先出模式

// DeferFunc testing
func DeferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i = %d \n", i)
	}
}
