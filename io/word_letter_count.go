package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func counter(input string) {
	nrchars += len(input) - 2 // -2 for \r\n
	nrwords += len(strings.Fields(input))
	nrlines++
}

func main() {
	nrchars, nrwords, nrlines = 0, 0, 0
	// inputReader 是一个指向 bufio.Reader 的指针
	inputReader := bufio.NewReader(os.Stdin) // 绑定标准输入
	fmt.Println("Please enter some input, type S to stop: ")

	for {
		input, err := inputReader.ReadString('\n') // 此处使用单引号 '' 转换类型为 byte
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		if input == "S\r\n" { // on Linux it is "S\n"
			fmt.Println("Hear are the counts: ")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		counter(input)
	}
}
