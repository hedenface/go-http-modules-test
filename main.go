package main

import (
	"flag"
	server "github.com/hedenface/go-http-modules-test/server"
)

func main() {

	var config string
	flag.StringVar(&config, "config", "", "the location of the main config file")
	flag.Parse()

	server.RegisterModules(server.ReadConfig(config))

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