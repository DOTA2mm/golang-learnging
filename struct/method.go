package main

import (
	"fmt"
)

// 定义一个 slice 类型
// 此处必须用 type 关键字申明
// 否则 intVecter is not a type
type intVecter []int

// 非 struct 类型的 method
func (v intVecter) sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func init() {
	fmt.Println("Sum is ", intVecter{1, 3, 4}.sum())
}
