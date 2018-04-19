package main

import (
	"fmt"
)

func main() {
	var a [10]int
	a[0] = 1
	a[1] = 2
	a[3] = 3
	fmt.Println(a)

	testArray()

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice = ", slice)
	slice1 := make([]int, 5)
	fmt.Println("slice1 = ", slice1)

	slice2 := make([]int, 0, 5)
	printSlice("slice2", slice2)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func testArray() {
	p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(p)
	for i := 0; i < 10; i++ {
		fmt.Printf("p[%d] = %d\n", i, p[i])
	}

	fmt.Println("p[:3] == ", p[:3])
}
