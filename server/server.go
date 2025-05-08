package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

//  TODO : create either CLI tool or some user interface for server start

type option func(*TCPServer) error

type TCPServer struct {
	address  string
	listener net.Listener
	rw       *bufio.ReadWriter
	conns    map[net.Conn]struct{}
	mu       sync.Mutex
}

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

	s.conns = make(map[net.Conn]struct{})

	return s, nil
}

// Starts the TCPServer, listening for any connections given on the address
// given by the server's configuration
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

		go s.handleConnections(conn)
	}

}

func (s *TCPServer) broadcast(msg string) {

}

// TODO: handle multiple connections at once...
func (s *TCPServer) handleConnections(conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	s.rw = rw

	for {

		req, err := s.rw.ReadString('\n') // reads string until new line char
		if err != nil {
			s.rw.WriteString("FAILED TO READ INPUT")
			s.rw.Flush()
			return
		}

		s.rw.WriteString(fmt.Sprintf("Request received: %s", req))
		s.rw.Flush()

		if req == "/quit\n" {
			s.Stop()
		}

	}

}

// Stops the TCPServer from listening for any connections, closing the server.
// Depends on if the server is listening
func (s *TCPServer) Stop() {
	s.listener.Close()
}
