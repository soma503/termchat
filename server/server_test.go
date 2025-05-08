package server_test

import (
	"bytes"
	"net"
	"testing"
)

func TestNETServer_Request(t *testing.T) {
	tt := []struct {
		test    string
		address string
		payload []byte
		want    []byte
	}{
		{
			"Sending a simple request returns result",
			"localhost:9999",
			[]byte("hello world\n"),
			[]byte("Request received: hello world"),
		},
		{
			"Sending another simple request works",
			"localhost:9999",
			[]byte("goodbye world\n"),
			[]byte("Request received: goodbye world"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.test, func(t *testing.T) {
			conn, err := net.Dial("tcp", tc.address)
			if err != nil {
				t.Error("could not connect to TCP server: ", err)
			}
			defer conn.Close()

			if _, err := conn.Write(tc.payload); err != nil {
				t.Error("could not write payload to TCP server:", err)
			}

			out := make([]byte, 1024)
			if _, err := conn.Read(out); err == nil {
				if bytes.Equal(out, tc.want) {
					t.Error("response did match expected output")
				}
			} else {
				t.Error("could not read from connection")
			}
		})
	}
}

func TestTCPServer_HandleConnections(t *testing.T) {
	tt := []struct {
		test    string
		address string
		payload string
		want    string
	}{}

	for _, tc := range tt {

		t.Run(tc.test, func(t *testing.T) {
			conn, err := net.Dial("tcp", tc.address)

			if err != nil {
				t.Error("could not connect to TCP server: ", err)
			}
			defer conn.Close()

		})

	}
}
