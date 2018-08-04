// calculator.go
// 例如：计算 1+ 2 = 3 输入：1 回车 2 回车 + 回车 --> 输出 3
package main

import (
	"bufio"
	"fmt"
	"golang-learning/io/stack"
	"os"
	"strconv"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	calc := new(stack.Stack)
	fmt.Println("Give a number, an operator (+, -, * , /), or q to stop: ")

	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Input error !")
			os.Exit(1)
		}
		token = token[0 : len(token)-2] // remove \r\n

		switch {
		case token == "q":
			fmt.Println("Calculator stopped")
			return
		case token >= "0" && token <= "999999":
			i, _ := strconv.Atoi(token)
			calc.Push(i)
		case token == "+":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p+q)
		case token == "-":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p-q)
		case token == "*":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p*q)
		case token == "/":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p/q)
		default:
			fmt.Println("No valid input")
		}

	}
}
