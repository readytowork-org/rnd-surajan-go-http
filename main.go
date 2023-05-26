package main

import (
	"fmt"
	"rnd-surajan-go-http/env"

	// Import your lessons.
	// "rnd-surajan-go-http/hello_world"
	"rnd-surajan-go-http/http_server"
)

func main() {
	fmt.Println("Welcome to my R&D of 'net/http' package in Go.")
	// Initialize/Set environment variables.
	env.Init()
	// Call methods from packages.
	// hello_world.HelloWorld()
	http_server.HttpServer()
}
