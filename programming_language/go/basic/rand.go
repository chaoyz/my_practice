package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	rand.Seed(42)
	fmt.Println("rand seed:", rand.Intn(30))

	fmt.Println("output name:", math.Pi)
}
