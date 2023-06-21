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

// Client Tcp: bind, connect, read writer, close
func Client() {

	buffer := make([]byte, 1024)
	bind := fmt.Sprintf("%s:%s", IP, PORT)
	fmt.Println("Connecting in ", bind)
	// realiza bind e connect
	conn, err := net.Dial("tcp", bind)
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}

	for {

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(text) == "stop" {
			break
		}
		_, err = conn.Write([]byte(text + "\n"))
		if err != nil {
			log.Panicln(err)
		}
		b, err := conn.Read(buffer)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Printf("Received %d bytes from server that says: %s\n", b, string(buffer))

	}
}

func main() {
	Client()
}
