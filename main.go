package main

import "log"

func main() {
	s := NewServer()
	err := s.Listen("localhost", "8080")
	if err != nil {
		log.Println("Error listening")
	}
	defer s.Close()

	s.Start()
}
