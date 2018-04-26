package main

import (
	"fmt"
)

func main() {
	a := new(bool)
	fmt.Println("bool:", *a)

	b := new(int)
	fmt.Println("int:", *b)

	c := new(float32)
	fmt.Println("float32:", *c)

	d := new(string)
	fmt.Println("string:", *d)

	e := new([]int)
	fmt.Println("[]int:", *e)

	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
}
