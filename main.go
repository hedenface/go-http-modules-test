package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	server "github.com/hedenface/go-http-modules-test/server"
)

func main() {

	var config string
	flag.StringVar(&config, "config", "", "the location of the main config file")
	flag.Parse()

	server.RegisterModules(server.ReadConfig(config))

	http.HandleFunc("/", rootHandler)

	for key, pathFunc := range server.PathFuncs {
		path := fmt.Sprintf("%s", key)
		http.HandleFunc(path, pathFunc)
	}

	log.Fatal(http.ListenAndServe(":9090", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just servin from main over here...")
}