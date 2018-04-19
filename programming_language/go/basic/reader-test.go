package main

import (
	"fmt"
)

type test struct{}

func (t test) Reader(buff []byte) (int, error) {
	count := 0
	fmt.Println("buff len:", len(buff))
	for i := 0; i < len(buff); i++ {
		buff[i] = 'A'
		count++
	}
	return count, nil
}

func main() {
	t := test{}

	buff := make([]byte, 8)
	for i := 0; i < 100; i++ {
		n, err := t.Reader(buff)
		fmt.Printf("buff:%v n:%v err:%v", buff, n, err)
	}
}
