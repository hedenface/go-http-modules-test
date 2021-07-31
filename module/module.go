package main

import "fmt"

type helloWorld string

func (h helloWorld) HelloWorld() {
	fmt.Println("Hello World!")
}

var HelloWorld helloWorld
