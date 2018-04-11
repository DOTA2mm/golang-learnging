package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

// Color 声明自定义类型 Color (byte 类型的别名)
type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // a slice of boxes

// 定义 Box 的 method 算体积
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// 该 method 的 receiver 为 Box 类型的指针
func (b *Box) SetColor(c Color) {
	// 如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
	// *b.color = c
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.0
	k := Color(WHITE)
	// 遍历取体积最大的
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		// 如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method
		// (&bl[i]).SetColor(BLACK)
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	string := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return string[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	// len 函数取 slice 长度
	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}

// 同时 method 也支持 **继承** 与 **重写**
// 如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method
