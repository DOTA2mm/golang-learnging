package main

import (
	"fmt"
)

type list []int

func (l list) Len() int {
	return len(l)
}
func (l *list) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func init() {
	fmt.Println("======== methodset.go ========")
	// 值类型
	var lst list
	// CountInto(lst, 1, 10) // 报错：Append 定义在指针上
	if LongEnough(lst) {
		fmt.Printf(" - lst is long enough\n")
	} else {
		fmt.Printf(" - lst is not long enough\n")
	}

	// 指针类型
	plst := &list{}
	// or plst := new(list)
	CountInto(plst, 1, 10) // VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *list can be dereferenced for the receiver
		// 指针自动解引用
		fmt.Printf(" - plst is long enough\n")
	}
	fmt.Println("======== methodset.go ========")
}

// 在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：
// 	- 指针方法可以通过指针调用
// 	- 值方法可以通过值调用
// 	- 接收者是值的方法可以通过指针调用，因为指针会首先**被解引用**
// 	- 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
// 将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。
