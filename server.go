package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	// messages    []string
	clients []*Client
	// host     string
	// port     string
	listener net.Listener
	// mutex       *sync.Mutex
	countClient int
}

func NewServer() *Server {
	return &Server{
		clients: make([]*Client, 0, 20),
	}
}

func (s *Server) Listen(host, service string) error {
	fmt.Println("Server listen")
	l, err := net.Listen("tcp", host+":"+service)
	if err != nil {
		return err
	}
	s.listener = l
	fmt.Println("Server listening")
	return nil
}

func (s *Server) Close() {
	s.listener.Close()
}

func (s *Server) Start() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print(err)
		} else {
			client := NewClient(s.countClient, conn, s)
			s.countClient++
			s.clients = append(s.clients, client)
			go client.Start()
		}
	}
}

func (s *Server) SendAll(msg string, id int) {
	for i, value := range s.clients {
		if i != id {
			value.Send(msg)
		}
	}
}
