package main

import (
	"fmt"
)

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (stock stockPosition) getValue() float32 {
	return stock.sharePrice * stock.count
}

type car struct {
	make  string
	model string
	price float32
}

func (car car) getValue() float32 {
	return car.price
}

type valuable interface {
	getValue() float32
}

// 所有实现了 valuable 接口的类型都可调用该函数
func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func init() {
	fmt.Println("=========== valuable.go =========")
	//  直接定义 valuable 类型
	var o valuable = stockPosition{"GOOG", 578.19, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)

	// 等价于显式的类型转换
	st := stockPosition{"GOOG", 578.19, 4}
	showValue(valuable(st))
	car := car{"BMW", "M3", 66500}
	showValue(valuable(car))
	fmt.Println("=========== valuable.go =========")
}
