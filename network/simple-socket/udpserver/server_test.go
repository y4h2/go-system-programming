package main

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUDP(t *testing.T) {
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

		server := NewUDPServer(addr)
		go func() {
			server.Run()
		}()
		defer server.Close()

		t.Run(spec.description, func(t *testing.T) {

			conn, err := net.Dial("udp", addr)
			assert.NoError(err)

			_, err = conn.Write([]byte(spec.send))
			assert.NoError(err)

			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			assert.Greater(n, 0)
			assert.Equal(spec.receive, string(buf[:n]))

			conn.Close()

		})
	}

}
