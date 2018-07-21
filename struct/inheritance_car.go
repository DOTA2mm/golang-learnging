package main

import (
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

// define a behavior for Car
func (car Car) numberOfWheels() int {
	return car.wheelCount
}

type Mercedes struct {
	Car // anonymous field Car
}

func (m Mercedes) sayHi() {
	fmt.Println("Hi, I'm Mercedes")
}

func (_ Car) Start() {
	fmt.Println("Car is started")
}
func (_ Car) Stop() {
	fmt.Println("Car is stopped")
}

func (c Car) GoToWorkIn() {
	c.Start()
	c.Stop()
}

func init() {
	m := Mercedes{Car{4, nil}}
	fmt.Println("A Mercedes has this many wheels: ", m.numberOfWheels())
	m.GoToWorkIn()
	m.sayHi()
}
