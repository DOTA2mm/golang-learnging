package main

import "fmt"

func main() {
	// test side_effect function
	n := 0
	reply := &n
	multiply(10, 5, reply)
	fmt.Println("multiply: ", *reply)
	fmt.Println("multiply: ", n)

	// test varnumpar
	x := minimum(1, 3, 2, 0)
	fmt.Printf("The minimun is: %d\n", x)
	slice := []int{7, 9, 3, 2, 1}
	x = minimum(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x)

	// test fibonacci
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}

	// test factory function
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	fmt.Printf("file addBmp: %s\n", addBmp("file"))
	fmt.Printf("file addJpeg: %s\n", addJpeg("file"))
}
