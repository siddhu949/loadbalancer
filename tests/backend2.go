package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9002")
	if err != nil {
		fmt.Println("Error starting Backend 2:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Backend Server 2 running on 9002")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		fmt.Println("Request received on Backend 2")
		conn.Write([]byte("Response from Backend 2\n"))
		conn.Close()
	}
}
