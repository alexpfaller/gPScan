package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/gookit/color"
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

	if port == "-a" {
		allScan(80)
	} else if port == "-af" {
		allScan(5)
	} else if port == "-asf" {
		allScan(0)
	} else {
		normalScan()
	}
}
func normalScan() {
	portStr := dest + ":" + port
	_, err := net.Dial(mode, portStr)
	if err == nil {
		color.Danger.Printf("[-] Port %v is currently open!", port)
	} else {
		color.Blue.Printf("[+] Port %v is currently closed!", port)
	}
}

func allScan(s int) {
	color.Cyan.Printf("Press ctrl + c  to abort!\n")
	time.Sleep(4 * time.Second)
	openPorts := make([]string, 0)
	for i := 0; i <= 65535; i++ {
		p := fmt.Sprintf("%v", i)
		portStr := dest + ":" + p
		_, err := net.Dial(mode, portStr)
		if err == nil {
			color.Danger.Printf("[-] Port %v is open\n", p)
			openPorts = append(openPorts, p)
		} else {
			color.Blue.Printf("[+] Port %v is closed\n", p)
		}
		time.Sleep(time.Duration(s) * time.Millisecond)
	}
	color.Danger.Printf("Following ports are currently open:\n%v\n", openPorts)

}
