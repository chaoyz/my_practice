package main

import "fmt"

type student struct {
	id   int64
	name string
	age  int
}

func (s *student) printName() {
	fmt.Println("name:", s.name)
}

func (s *student) changeName(name string) {
	s.name = name
}

func main() {
	yc := student{
		1, "yangchao", 25,
	}
	yc.printName()
	yc.changeName("yc")
	yc.printName()
}
