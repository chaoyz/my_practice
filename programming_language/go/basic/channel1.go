package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := make(chan string, 4)
	for i := 0; i < 4; i++ {
		go put(c, strconv.Itoa(i))
	}
	for i := 0; i < 4; i++ {
		<-c
	}
	fmt.Println("exec done.")

	a := 1
	a++
	fmt.Println(a)
}

func put(c chan string, name string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v, %v\n", name, i)
	}
	c <- name
}
