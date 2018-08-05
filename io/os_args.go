package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	hello := "Hello "
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		hello += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(hello)
}
