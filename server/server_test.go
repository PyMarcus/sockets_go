package main

import (
	"fmt"
	"log"
	"net"
	"testing"
)

// testa a conexão
func TestServer(t *testing.T) {
	bind := fmt.Sprintf("%s:%s", IP, PORT)
	_, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalln(err)
	}
}
