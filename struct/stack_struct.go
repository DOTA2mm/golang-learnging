package main

import (
	"fmt"
	"strconv"
)

const LIMIT int = 5

type stack struct {
	index int
	data  [LIMIT]int
}

func (st *stack) Push(n int) {
	if st.index+1 > LIMIT {
		return
	}
	st.data[st.index] = n
	st.index++
}
func (st *stack) Pop() int {
	if st.index-1 < 0 {
		return -1
	}
	st.index--
	return st.data[st.index]
}
func (st stack) String() (str string) {
	for i := 0; i < st.index; i++ {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(st.data[i]) + "] "
	}
	return
}

func init() {
	st1 := new(stack)
	fmt.Printf("%v\n", st1)
	st1.Push(3)
	fmt.Printf("%v\n", st1)
	st1.Push(7)
	fmt.Printf("%v\n", st1)
	st1.Push(10)
	fmt.Printf("%v\n", st1)
	st1.Push(99)
	fmt.Printf("%v\n", st1)
	p := st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
}
