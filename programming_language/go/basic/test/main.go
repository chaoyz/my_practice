package main

import (
	"fmt"

	"./class"
)

func (test *class.Class) Print() {
	fmt.Println("test:", test.i, test.j)
}

func main() {
	var t = &Test{0, 1}
	t.print()
}
