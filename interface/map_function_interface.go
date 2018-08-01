package main

import (
	"fmt"
)

type obj interface{}

func mapFunc(mf func(obj) obj, list []obj) []obj {
	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}

func main() {
	mf := func(i obj) obj {
		switch i.(type) {
		case int:
			// i.(type) 类型转换？
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}

	is1 := []obj{0, 1, 2, 3, 4, 5}
	res1 := mapFunc(mf, is1)
	for _, v := range res1 {
		fmt.Println(v)
	}
	println()
	ss1 := []obj{"a", "b", "c", "d", "e", "f"}
	res2 := mapFunc(mf, ss1)
	for _, v := range res2 {
		fmt.Println(v)
	}
}
