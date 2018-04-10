package main

import (
	"fmt"
)

// struct 可供声明创建自定义类型
// struct 有 4 种声明使用方式

// 声明一个新的类型, 拥有两个字段为 name, age
type person struct {
	name string
	age  int
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是 **传值** 的
func older(p1, p2 person) (person, int) {
	if p1.age > p2.age { // 比较年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func init() {
	var tom person
	// struct 的多种赋值方式
	// 赋值初始化
	tom.name, tom.age = "Tom", 18
	// 通过field:value的方式初始化，这样可以任意顺序
	bob := person{age: 25, name: "Bob"}
	// 按照struct定义顺序初始化值
	paul := person{"Paul", 28}

	tbOlder, tbDiff := older(tom, bob)
	tpOlder, tpDiff := older(tom, paul)
	bpOlder, bpDiff := older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tbOlder.name, tbDiff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tpOlder.name, tpDiff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bpOlder.name, bpDiff)
}
