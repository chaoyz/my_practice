package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func nd(z float32) {
	for i := 1; i <= 10; i++ {
		fmt.Println(sqrt(z))
		z = z - (z*z - x)
	}
}

func main() {
	// fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 2, 20),
	)
}
