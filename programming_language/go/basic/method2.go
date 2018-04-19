package main

import (
	"fmt"
)

type num float64

func (param num) print() {
	fmt.Println("param:", param)
}

func main() {
	i := num(100)
	i.print()
}
