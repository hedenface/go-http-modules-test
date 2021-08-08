package main

import (
	"fmt"
	server "github.com/hedenface/go-http-modules-test/server"
)

func Register() {
	server.RegisterCallback(server.BeforePrinting, BeforePrinting)
	server.RegisterCallback(server.AfterPrinting, AfterPrinting)
}

func BeforePrinting(hook server.Hook) {
	if hook != server.BeforePrinting {
		return
	}

	fmt.Println("THE MODULE IS PRINTING BEFORE")
}

func AfterPrinting(hook server.Hook) {
	if hook != server.AfterPrinting {
		return
	}

	fmt.Println("THE MODULE IS PRINTING AFTER")
}