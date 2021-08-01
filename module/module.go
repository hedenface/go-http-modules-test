package main

import "fmt"

type helloWorld struct {
	s string
}

func (h helloWorld) HelloWorld() {
	fmt.Println(h.s)
}

var HelloWorld helloWorld
