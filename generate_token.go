package main

import (
	"fmt"
	"loadbalancer/internal/firewall"
)

func main() {
	token, err := firewall.GenerateJWT()
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	fmt.Println("Generated JWT Token:", token)
}
