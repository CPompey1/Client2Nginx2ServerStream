package main

import (
	"fmt"
	"net"
)

func main() {
	// Start listening on localhost:9000
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on localhost:9000")

	for {
		// Wait for a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			break
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		if n == 0 {
			fmt.Println("Connection closed by client")
			return
		}

		// Print the received message
		fmt.Printf("Received message: %s", string(buf[:n]))
	}
}
