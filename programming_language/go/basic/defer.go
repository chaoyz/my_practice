package main

import (
	"fmt"
)

func main() {
	var i int = 0
	defer fmt.Println("i:", i)
	fmt.Println("h")
	i++
	fmt.Println("i`:", i)
}
