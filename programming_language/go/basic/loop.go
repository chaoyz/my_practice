package main

import (
	"fmt"
)

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forLoop1() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func forLoop2() {

	sum := 1

	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	forLoop2()
}
