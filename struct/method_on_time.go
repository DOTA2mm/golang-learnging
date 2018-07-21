package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func init() {
	mt := myTime{time.Now()}
	// 调用匿名Time上的String 方法
	fmt.Println("Full time now: ", mt.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars: ", mt.first3Chars())
}
