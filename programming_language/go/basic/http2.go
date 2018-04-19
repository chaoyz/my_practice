package main

import (
	"fmt"
	"log"
	"net/http"
)

type handle struct {
	Greeting string
	Punt     string
	Who      string
}

func (handler *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(handler.Greeting)
	s := fmt.Sprint(handler.Greeting, handler.Punt, handler.Who)
	fmt.Fprint(w, s)
}

func main() {
	http.Handle("/struct", &handle{"Hello", ":", "yc!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
