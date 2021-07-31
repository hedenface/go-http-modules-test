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

type HelloWorld interface {
	HelloWorld()
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

	var helloWorld HelloWorld
	helloWorld, ok := symHelloWorld.(HelloWorld)
	if !ok {
		fmt.Println("unexpected type from module")
		os.Exit(1)
	}

	helloWorld.HelloWorld()
}
