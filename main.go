package main

import (
	"flag"
	"fmt"
	server "github.com/hedenface/go-http-modules-test/server"
)

func main() {

	var config string
	flag.StringVar(&config, "config", "", "the location of the main config file")
	flag.Parse()

	server.RegisterModules(server.ReadConfig(config))

	fmt.Println("Before printing:")
	for _, hookFunc := range server.HookFuncs[server.BeforePrinting] {
		hookFunc(server.BeforePrinting)
	}

	fmt.Println("Printing:")
	fmt.Println("Hello!")

	fmt.Println("After printing:")
	for _, hookFunc := range server.HookFuncs[server.AfterPrinting] {
		hookFunc(server.AfterPrinting)
	}
}
