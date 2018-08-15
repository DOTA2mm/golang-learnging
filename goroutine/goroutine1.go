package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "Hello"
	ch <- "Golang"
	ch <- "Nice"
	ch <- "Programing"
	ch <- "Language"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func init() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	// time.Sleep(1e9)
}
