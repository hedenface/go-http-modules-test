package main

import (
	"fmt"
	"plugin"
	"flag"
	"os"
	"bufio"
	"io"
)

// type helloWorld struct {
// 	s string
// 	HelloWorld func()
// }

func readConfig(config string) {
	file, err := os.Open(config)
	if err != nil {
		fmt.Println("Unable to open file:", config)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		module_path := fmt.Sprintf("%s", line)

		module, err := plugin.Open(module_path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		symRegister, err := module.Lookup("Register")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)			
		}

		symRegister.(func())()
	}
}

func main() {

	var config string
	flag.StringVar(&config, "config", "", "the location of the main config file")
	flag.Parse()

	readConfig(config)

	// mod, err := plugin.Open("./module/module.so")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// symHelloWorld, err := mod.Lookup("HelloWorld")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// h := helloWorld{"hello there"}
	// h, ok := symHelloWorld.(HelloWorld)
	// if !ok {
	// 	fmt.Println("unexpected type from module")
	// 	os.Exit(1)
	// }

	// h.HelloWorld()
}
