package server_test

import (
	"github.com/soma503/termchat/server"
	"testing"
)

func TestServer_ServerCreation(t *testing.T) {
	t.Parallel()
	_, err := server.NewServer()
	if err != nil {
		t.Error("err with new server")
	}
}

func TestServer_ServerCreationWithArgs(t *testing.T) {
	t.Parallel()

	_, err := server.NewServer(
		server.WithHost("localhost"),
		server.WithPort("3333"),
	)

	if err != nil {
		t.Error(" err with new server with arguments ")
	}

}
