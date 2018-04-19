package main

import "fmt"

func main() {
	var a = make([]int, 0, 5)

	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(cap(a))

	a = append(a, 2)
	fmt.Println(a)
	fmt.Println(cap(a))
}
