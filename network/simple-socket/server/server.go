package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"
)

type Server struct {
	address string
	server  net.Listener
}

func (s *Server) Run() error {
	server, err := net.Listen("tcp", s.address)
	if err != nil {
		fmt.Printf("Fail to start server, %s\n", err)
		return errors.Wrap(err, "Fail to start server")
	}
	s.server = server

	fmt.Println("Server Started ...")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect, %s\n", err)
			break
		}
		go connHandler(conn)
	}

	return nil
}

func (s *Server) Close() {
	s.server.Close()
}

func NewServer(address string) *Server {
	return &Server{address: address}
}

func connHandler(c net.Conn) {
	if c == nil {
		return
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		inputs := strings.Split(inStr, " ")
		switch inputs[0] {
		case "ping":
			c.Write([]byte("pong\n"))
		case "echo":
			echoStr := strings.Join(inputs[1:], " ") + "\n"
			c.Write([]byte(echoStr))
		case "quit":
			c.Close()
			break
		default:
			c.Write([]byte("Unsupported command \n"))
			// fmt.Printf("Unsupported command: %s\n", inputs[0])
		}
	}
	fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
}

func main() {
	NewServer(":1208")
}
