package main

import (
	"fmt"
	"net"
	"os"
)

var (
	mode = os.Args[1]
	dest = os.Args[2]
	port = os.Args[3]
)

func main() {
	if mode == "tcp" || mode == "udp" {
		normalScan()
	} else {
		fmt.Println("Please enter a supported protocol! (tcp or udp)")
	}
}
func normalScan() {
	portStr := dest + ":" + port
	_, err := net.Dial(mode, portStr)
	if err == nil {
		fmt.Printf("Port %v is currently open!", port)
	} else {
		fmt.Printf("Port %v is currently closed!", port)
	}
}
