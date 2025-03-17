package server

import (
	"net"
)

type server struct {
	port     string
	host     string
	listener net.Listener
}

type option func(*server) error

func WithPort(port string) option {
	return func(s *server) error {
		s.port = port
		return nil
	}
}

func WithHost(host string) option {
	return func(s *server) error {
		s.host = host
		return nil
	}
}

func NewServer(opts ...option) (*server, error) {

	s := &server{}

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (s *server) Start() error {
	listener, err := net.Listen(s.host, s.port)

	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		HandleConnections(conn)

	}
}

func HandleConnections(conn net.Conn) {

}
