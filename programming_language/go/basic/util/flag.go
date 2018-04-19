package main

import (
	"flag"
	"fmt"
	"os"
)

var commandLine = flag.Bool("n", false, "don`t print final newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.Parse()
	s := ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
		fmt.Println(s)
	}
	fmt.Println(*commandLine)
	if !*commandLine {
		s += Newline
	}
	os.Stdout.WriteString(s)
}
