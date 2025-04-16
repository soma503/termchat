package server

import (
	"fmt"
	"net"
	"strings"
)

type option func(*TCPServer) error

func WithAddress(address string) option {
	return func(s *TCPServer) error {
		s.address = strings.ToLower(address)
		return nil
	}
}

// Creates a new TCPServer, using the options given in order to configure it.
// Returns any errors in the configuring process
func NewServer(opts ...option) (*TCPServer, error) {

	s := &TCPServer{}

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

type TCPServer struct {
	address  string
	listener net.Listener
}

// Starts the
func (s *TCPServer) Run() {
	var err error
	s.listener, err = net.Listen("tcp", s.address)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	defer s.listener.Close()

	fmt.Println("listening on address: ", s.address)

	for {

		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		go handleConnections(conn)
	}

}

func handleConnections(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error handling connection: ", err)
			return
		}

		s := buffer[:n]
		//TODO: add some way clients can stop the freaking server :(

		fmt.Printf("received: %s\n", s)

	}

}

// Stops the TCPServer from listening for any connections, closing the server.
// Depends on if the server is listening
func (s *TCPServer) Stop() {
	s.listener.Close()
}
