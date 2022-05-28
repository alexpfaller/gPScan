package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var (
		mode    = os.Args[1]
		port    = os.Args[2]
		portStr = "localhost:" + port
	)

	_, err := net.Dial(mode, portStr)
	if err == nil {
		fmt.Printf("Port %v is currently open!", port)
	} else {
		fmt.Printf("Port %v is currently closed!", port)
	}
}
