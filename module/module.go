package main

import (
	"fmt"
	"net/http"
	server "github.com/hedenface/go-http-modules-test/server"
)

func Register() {
	server.RegisterCallback("/module", ModuleServe)
}

func ModuleServe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there, this is your captain, the module speaking")
}
