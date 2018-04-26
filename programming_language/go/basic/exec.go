package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("date")
	// var out bytes.Buffer
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// fmt.Println(out.String())
}
