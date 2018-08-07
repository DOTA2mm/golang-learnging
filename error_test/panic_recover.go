package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		// recover 只能在 defer 修饰的函数中使用：用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil，且没有其它效果。
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call \n")
}

func main() {
	fmt.Printf("Calling test\n")
	test()
	fmt.Printf("Test completed\n")
}
