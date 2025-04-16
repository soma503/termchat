package main

import (
	"fmt"
	"github.com/soma503/termchat/server"
	"os"
)

func main() {

	s, err := server.NewServer(
		server.WithAddress("localhost:9999"),
	)

	if err != nil {
		fmt.Println("ERROR: Could not create server")
		os.Exit(1)
	}

	fmt.Println("server creation was successful")

	s.Run()

}
