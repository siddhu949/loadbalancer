package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9003")
	if err != nil {
		fmt.Println("Error starting Backend 3:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Backend Server 3 running on 9003")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		fmt.Println("Request received on Backend 3")
		conn.Write([]byte("Response from Backend 3\n"))
		conn.Close()
	}
}
