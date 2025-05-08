package main

import (
	"fmt"
	"github.com/soma503/termchat/client"
)

func main() {
	fmt.Println("OPENING CLIENT...")
	c, _ := client.NewClient(
		client.WithAddress("localhost:9999"),
	)
	c.Start()
}
