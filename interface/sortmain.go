package main

import (
	"fmt"
	"golang-learning/interface/sort"
)

func ints() {
	data := []int{35, 36, 473, -2, 345, 0, 1, 4235, 45, -23, 1}
	a := sort.IntArray(data) // 转换为 IntArray 类型
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

type day struct {
	num       int
	shortName string
	longName  string
}
type dayArray []*day

// type dayArray struct{ data []day }

func (p dayArray) Len() int           { return len(p) }
func (p dayArray) Less(i, j int) bool { return p[i].num < p[j].num }
func (p dayArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := dayArray{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}

	sort.Sort(data)
	if !sort.IsSorted(data) {
		panic("fail")
	}
	for i, d := range data {
		fmt.Printf("The %d is %s \n", i, d.longName)
	}
}

func init() {
	ints()
	days()
}
