package main

import (
	"fmt"
)

func init() {
	for i := 0; i < 4; i++ {
		// 匿名函数可以分配不同的内存地址
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
