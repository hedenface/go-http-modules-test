package main

import (
	"fmt"
	"plugin"
	"os"
)

// func handleError(err error) {
// 	fmt.Println(err)
// 	os.Exit(1)
// }

// type HelloWorld interface {
// 	HelloWorld()
// }

// type helloWorld interface {
// 	s string
// 	HelloWorld() int
// }

type helloWorld struct {
	s string
	HelloWorld func()
}

func main() {

	mod, err := plugin.Open("./module/module.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symHelloWorld, err := mod.Lookup("HelloWorld")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	h := helloWorld{"hello there"}
	h, ok := symHelloWorld.(HelloWorld)
	if !ok {
		fmt.Println("unexpected type from module")
		os.Exit(1)
	}

	h.HelloWorld()
}
