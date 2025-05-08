package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type option func(c *Client) error

type Client struct {
	address  string
	username string
}

func WithAddress(address string) option {
	return func(c *Client) error {
		c.address = strings.ToLower(address)
		return nil
	}
}

func NewClient(opts ...option) (*Client, error) {

	s := &Client{}

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (c *Client) Start() {

	conn, err := net.Dial("tcp", c.address)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	defer conn.Close()

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text := sc.Text()
		payload := []byte(text + "\n")

		if _, err = conn.Write(payload); err != nil {
			fmt.Println("Error: ", err)
			return
		}

		out := make([]byte, 1024)
		if _, err := conn.Read(out); err == nil {
			fmt.Printf("SERVER SAID: %s\n", string(out))
		} else {
			fmt.Println(err)
			return
		}

		if text == "/quit" {
			fmt.Println("QUITTING")
			return
		}
	}

}

func handleConnection(conn net.Conn) {

}
