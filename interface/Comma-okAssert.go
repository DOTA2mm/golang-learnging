package main

// 通过 Comma-ok 断言来判断 interface 变量里存储的到底是何种类型的对象

import (
	"fmt"
	"strconv"
)

// 定义一个空的 interface
type Element interface{}

// 定义一个存放 Element 类型的 slice 类型
type List []Element

type Person struct {
	name string
	age  int
}

// 定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func init() {
	// slice list 可以存任意类型值
	list := make(List, 3)
	list[0] = 1       // int
	list[1] = "Hello" // string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		// Comma-ok 断言
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

	// or 使用 switc...case 改写如下
	// `element.(type)`语法不能在switch外的任何逻辑里面使用，
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
