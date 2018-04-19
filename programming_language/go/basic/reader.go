package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	s := strings.NewReader("sadhfjk123ncsv09u24ntds0-wn1`90adsun2bvfjoprt")

	buff := make([]byte, 8)
	for {
		tmp, err := s.Read(buff)
		fmt.Printf("tmp = %v err = %v buff = %v\n", tmp, err, buff)
		fmt.Printf("buff[:tmp] = %q\n", buff[:tmp])
		if err == io.EOF {
			break
		}
	}
}
