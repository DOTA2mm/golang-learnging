package main

// multiply 改变外部变量（outside variable）
func multiply(a, b int, reply *int) {
	*reply = a * b
}
