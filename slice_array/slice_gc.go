package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// 切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。
// 只有在没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存。

var digistRegexp = regexp.MustCompile("[0-9]+")

func findDigists(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	// 只要该返回的切片不被释放，垃圾回收器就不能释放整个文件所占用的内存
	// return digistRegexp.Find(b)
	b = digistRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// 找到所有的
func findFileDigists(filename string) []byte {
	fileBytes, _ := ioutil.ReadFile(filename)
	b := digistRegexp.FindAll(fileBytes, len(fileBytes))
	c := make([]byte, 0)
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	return c
}

func init() {
	b := findDigists("testFile.md")
	fmt.Printf("Find digist is %v, convert to string is %s\n", b, string(b))

	b = findFileDigists("testFile.md")
	fmt.Printf("Find digist is %v, convert to string is %s\n", b, string(b))
	fmt.Println("-------------------------")
}
