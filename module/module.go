package main

import (
	"fmt"
	server "github.com/hedenface/go-http-modules-test/server"
)

func Register() {
	fmt.Println("hello")
	server.Hello()
}
