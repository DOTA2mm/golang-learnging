package main

// "A method is a function with an implicit first argument, called a receiver."
// method 的语法如下：
// func (r ReceiverType) funcName(parameters) (results)

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 定义 method
// ReceiverType 为 rectangle
func (r rectangle) area() float64 {
	return r.width * r.height
}

// 另一个 method
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func init() {
	r1 := rectangle{12, 2}
	r2 := rectangle{9, 4}
	c1 := circle{10}
	c2 := circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

// 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
// method里面可以访问接收者的字段
// 调用method通过.访问，就像struct里面访问字段一样
