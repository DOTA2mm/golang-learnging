package main

import (
	"fmt"
	"golang-learning/struct/struct_pack"
)

type Skills []string
type Human struct {
	name   string
	age    int
	weight int
}
type Student struct {
	Human      // 匿名字段，struct
	Skills     // 匿名字段 ， 自定义类型 string slice
	int        // 内置类型作为匿名字段
	speciality string
}

func main() {
	// 初始化学生 Jane
	jane := Student{Human: Human{"Jane", 25, 120}, speciality: "Biology"}

	// 访问相应字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)

	// 修改 skill 字段
	jane.Skills = []string{"anatomoy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)

	// 引用外部包
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.
	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)
}
