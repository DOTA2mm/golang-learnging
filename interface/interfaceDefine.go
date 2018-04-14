package main

// interface是一组method签名的组合，我们通过interface来定义对象的一组行为。

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human   // 匿名字段
	company string
	money   float32
}

// Human 实现 SayHi 方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I'm %s you can call me on %s\n", h.name, h.phone)
}

// Human 实现 Sing 方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Employee 重载 Human 的 SayHi 方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I'm %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Iinterface Men 被 Human, Student和 Employee 实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "111-222-xxx"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "222-222-xxx"}, "Harvard", 100}
	san := Student{Human{"San", 35, "111-333-xxx"}, "Google Inc", 2000}
	tom := Student{Human{"Tom", 29, "222-333-xxx"}, "LinkIn Inc", 1800}

	// 定义一个 Men 类型的变量 i
	// 因为 Human, Student和 Employee 都实现了 Men
	// 所以 i 可以存储这三种类型
	var i Men
	// i 能存储 Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// i 也能存储 Employee
	i = tom
	fmt.Println("This is Tom,  an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// 定义了 slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	// 这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, san, mike

	for _, value := range x {
		value.SayHi()
	}
}

// interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现
// Go通过interface实现了duck-typing:
// 即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子"
