package main

import (
	"fmt"
	"htmx/internal/server"
)

func main() {

	println("CARL!")
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
