package server

import (
	"fmt"
	"plugin"
	"os"
	"bufio"
	"io"
)

// type helloWorld struct {
// 	s string
// 	HelloWorld func()
// }

func ReadConfig(config string) (syms []plugin.Symbol) {
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

		syms = append(syms, symRegister)
	}

	return
}

func RegisterModules(syms []plugin.Symbol) {

	for _, sym := range syms {
		sym.(func())()
	}
}


func Hello() {
	fmt.Println("testing calling server func from module")
}