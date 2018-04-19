package main

import "fmt"

var a, b int = 0, 1

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	return func() int {
		a, b = b, (a + b)
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
