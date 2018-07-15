package main

import (
	"fmt"
)

func mapFunc(mf func(int) int, list []int) []int {
	result := make([]int, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}
	// for ix := 0; ix < len(list); ix++ {
	// 	retult[ix] = mf(list[ix])
	// }
	return result
}

func init() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mf := func(i int) int {
		return i * 10
	}
	fmt.Printf("result is %v\n", mapFunc(mf, list))
}
