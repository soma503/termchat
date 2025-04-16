package client

import (
	"fmt"
	"net"
)

type Client struct {
	port     string
	host     string
	username string
}

func Start() {

	conn, err := net.Dial("tcp", "localhost:9999")

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	defer conn.Close()

	data := []byte("waddup, king :p")

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}
