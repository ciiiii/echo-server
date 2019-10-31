package main

import (
	"fmt"
	"net"
	"os"
)

const (
	host = "0.0.0.0"
	port = 3333
)

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		return
	}
	defer l.Close()
	fmt.Printf("Listening on %s:%d\n", host, port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}
	version := os.Getenv("version")
	conn.Write([]byte(fmt.Sprintf("Message received at %s.", version)))
	fmt.Println(string(buf))
}
