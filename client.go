package main

import (
	"bufio"
	"log"
	"net"
	"sync"
)

type Client struct {
	id      int
	conn    net.Conn
	server  *Server
	mutex   *sync.Mutex
	message chan string
}

func NewClient(id int, conn net.Conn, s *Server) *Client {
	return &Client{
		id:      id,
		conn:    conn,
		server:  s,
		message: make(chan string),
	}
}

func NewClient1(id int, conn net.Conn, s *Server, c chan string) *Client {
	return &Client{
		id:      id,
		conn:    conn,
		server:  s,
		message: c,
	}
}

func (c *Client) Start() {
	for {
		userInput, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			log.Println("Error reading user input")
		}

		c.server.SendAll(userInput, c.id)
	}
}

func (c *Client) Send(msg string) {
	realSize := len(msg)
	c.mutex.Lock()
	size, err := c.conn.Write([]byte(msg))
	c.mutex.Unlock()

	if err != nil {
		log.Print(err)
	}

	if realSize > size {
		log.Printf("Отправлено %v из %v", realSize, size)
	}
}
