package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/gookit/color"
)

var Protocols = map[string]struct{}{
	"tcp": {},
	"udp": {},
}

func main() {

	all := flag.Bool("a", false, "Scan all ports")
	timeout := flag.Duration("timeout", time.Millisecond*200, "Dialing timeout")

	flag.Parse()
	args := flag.Args()

	if l := len(args); l < 2 || l > 3 {
		fmt.Printf("Wrong count of arguments, expected 2-3, got %d\n", l)
		flag.Usage()

		return
	}

	var (
		protocol = args[0]
		dest     = args[1]
		port     string
	)

	if len(args) == 3 {
		if !*all {
			port = args[2]

			_, err := strconv.Atoi(port)
			if err != nil {
				fmt.Println(fmt.Errorf("wrong port format: %w", err))

				return
			}
		} else {
			fmt.Println("'-a' flag is set, ignoring port")
		}

	}

	if _, ok := Protocols[protocol]; !ok {
		fmt.Printf("protocol '%s' not supported\n", protocol)
		fmt.Print("available: ")
		for k := range Protocols {
			fmt.Print(k, "")
		}
	}

	if !*all {
		if ScanPort(dest+":"+port, protocol, *timeout) {
			color.Green.Println("Port is open")
		} else {
			color.Danger.Println("Port is closed")
		}

		return
	}

	var open []int
	for i := 0; i <= 65535; i++ {
		if ScanPort(dest+":"+strconv.Itoa(i), protocol, *timeout) {
			open = append(open, i)
		}
	}

	fmt.Print("Open ports: ")
	color.Green.Println(open)
}

func ScanPort(address, mode string, timeout time.Duration) bool {
	conn, err := net.DialTimeout(mode, address, timeout)
	if err != nil {
		return false
	}

	err = conn.Close()
	if err != nil {
		fmt.Println(fmt.Errorf("close connection (%s): %w", address, err))
		return true
	}

	return true
}
