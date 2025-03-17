package main

import (
	"fmt"
	"github.com/soma503/termchat/server"
	"os"
)

func main() {

	s, err := server.NewServer(
		server.WithHost("localhost"),
		server.WithPort("9999"),
	)

	if err != nil {
		fmt.Println("ERROR: Could not create server")
		os.Exit(1)
	}

	fmt.Println("server creation was successful")

	s.Start()

}
