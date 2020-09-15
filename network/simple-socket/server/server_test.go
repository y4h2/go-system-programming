package main

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	specs := []struct {
		description string
		send        string
		receive     string
	}{
		{
			description: "test ping",
			send:        "ping",
			receive:     "pong",
		},
	}

	for _, spec := range specs {
		assert := assert.New(t)
		addr := ":1123"

		server := NewServer(addr)
		go func() {
			server.Run()
		}()
		defer server.Close()

		t.Run(spec.description, func(t *testing.T) {

			conn, err := net.Dial("tcp", addr)
			assert.NoError(err)

			_, err = conn.Write([]byte(spec.send))
			assert.NoError(err)

			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			assert.Greater(n, 0)
			assert.Equal(string(buf[:n]), spec.receive)

			conn.Close()

		})
	}

}

// test without start real server and client
// could be used to test TCP server and client internal logic
func TestPipe(t *testing.T) {
	assert := assert.New(t)

	server, client := net.Pipe()
	go func() {
		// Do some stuff
		server.Write([]byte("haha"))
		server.Close()
	}()

	// Do some stuff
	buf := make([]byte, 4096)
	n, err := client.Read(buf)
	assert.NoError(err)
	assert.Equal(string(buf[:n]), "haha")
	client.Close()
}
