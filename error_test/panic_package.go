package main

import (
	"fmt"
	"golang-learning/error_test/parse"
)

func main() {
	var examples = []string{
		"1 2 3 5 7",
		"10 20 3.14 233",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n ", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println(err) // 调用 ParseError 上的 String() 方法
			continue
		}
		fmt.Println(nums)
	}
}
