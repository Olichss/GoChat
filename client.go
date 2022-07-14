package main

import (
	"bufio"
	"log"
	"net"
	"sync"
)

type Client struct {
	name   string
	id     int
	conn   net.Conn
	server *Server
	mutex  *sync.Mutex
}

func NewClient(id int, conn net.Conn, s *Server, mtx *sync.Mutex) *Client {
	return &Client{
		id:     id,
		conn:   conn,
		server: s,
		mutex:  mtx,
	}
}

func (c *Client) Start() error {
	c.Send("Insert your name\n")
	userInput, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		log.Println("Error reading user input")
		return err
	}
	c.name = userInput[0 : len(userInput)-2]
	msg := Message{c.name, " присоединился к чату"}
	c.server.SendAll(msg.String(), c.id)
	return c.Reading()
}

func (c *Client) Reading() error {
	for {
		userInput, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			log.Println("Error reading user input")
			return err
		}
		msg := Message{c.name, userInput[0 : len(userInput)-2]}
		c.server.SendAll(msg.String(), c.id)
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
