package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type UDPServer struct {
	address string
	server  *net.UDPConn
}

func (s *UDPServer) Run() error {
	laddr, err := net.ResolveUDPAddr("udp", s.address)
	if err != nil {
		return errors.New("could not resolve UDP addr")
	}

	s.server, err = net.ListenUDP("udp", laddr)
	if err != nil {
		return errors.New("could not listen on UDP")
	}

	for {
		buf := make([]byte, 2048)
		n, conn, err := s.server.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			return err
		}
		if conn == nil {
			continue
		}

		go s.handleConnection(conn, buf[:n])
	}
}

func (u *UDPServer) handleConnection(addr *net.UDPAddr, cmd []byte) {
	u.server.WriteToUDP([]byte(fmt.Sprintf("Request recieved: %s", cmd)), addr)
}

func (s *UDPServer) Close() {
	s.server.Close()
}

func NewUDPServer(address string) *UDPServer {
	return &UDPServer{address: address}
}
