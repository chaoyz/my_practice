package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i, v := range a {
		fmt.Printf("index:%d, value:%d\n", i, v)
	}

	for i := range a {
		fmt.Printf("index:%d\n", i)
	}

	for _, v := range a {
		fmt.Printf("value:%d\n", v)
	}
}
