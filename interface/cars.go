package main

import "fmt"

type Any interface{}
type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}
type Cars []*Car

// 用给定的函数 f 处理 cars
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// 查找符合给定规则的 cars
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make(Cars, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

// Process cars and create new data.
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	// Prepare maps of sorted cars.
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make(Cars, 0)
	}
	sortedCars["Default"] = make(Cars, 0)

	// prepare appender function.
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}

	return appender, sortedCars
}

func main() {
	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	byd := &Car{"BYD 宋", "BYD", 2010}
	// query:
	allCars := Cars([]*Car{ford, bmw, merc, bmw2, byd})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	allBuildYear := allCars.Map(func(car *Car) Any {
		return car.BuildYear
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)
	fmt.Println("AllBuildYear: ", allBuildYear)
	//
	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}
