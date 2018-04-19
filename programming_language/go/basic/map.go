package main

import (
	"fmt"
)

type test struct {
	i, j int
}

func main() {
	m := make(map[string]test)
	m["test"] = test{
		1, 2,
	}
	fmt.Println(m)

	m1 := map[string]test{
		"1": test{
			1, 2,
		},
		"2": test{
			2, 3,
		},
	}
	fmt.Println(m1)

	fmt.Println(m1["test"])
}
