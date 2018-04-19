package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 针对结构体定义方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) print() {
	fmt.Println("X:", v.X)
	fmt.Println("Y:", v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	v.print()
}
