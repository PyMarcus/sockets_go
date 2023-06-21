package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	IP   string = "127.0.0.1"
	PORT string = "8808"
)

// Server TCP SYSCALLS-> bind(), Listen(), Accept(), write(), read(), close()
func Server() {
	fmt.Println("Running server...")
	bind := fmt.Sprintf("%s:%s", IP, PORT)

	// define o bind e listen
	listener, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Listening...")

	for {
		// chamada de sistema para aceitar multiplas conexões, bloqueando até receber
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Received connection from ", conn.RemoteAddr())
		}
		// cria goroutine (similar a uma thread em outras linguagens)
		go handleConnection(conn)
	}
}

// trata novas conexões
func handleConnection(conn net.Conn) {

	for {
		// le dados recebidos

		buffer, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		} else {
			fmt.Println("Client says: ", strings.Split(strings.TrimSpace(string(buffer)), "\n")[0])
		}

		// escrevendo no socket:
		_, err = conn.Write([]byte("Received successfully"))
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Sended message!")
		}

		if strings.Split(strings.TrimSpace(string(buffer)), "\n")[0] == "stop" {
			fmt.Println("Bye bye ", conn.RemoteAddr())
			os.Exit(0)
			break
		}
	}
}

func main() {
	Server()
}
