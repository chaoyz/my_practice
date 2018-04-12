package main

import "fmt"

type test struct {
	i int
	j int
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

// func main() {
// 	t := test{1, 2}
// 	t.i = 100
// 	fmt.Println(t.i)
// }
func main() {
	fmt.Println(v1, p, v2, v3)
}
