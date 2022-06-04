package gpscan

import (
	"flag"
	"fmt"
	"github.com/gookit/color"
	"net"
	"os"
	"strconv"
	"time"
)

const totalNumberOfPorts = 65535

type gPScan struct {
	protocol string
	timeout  time.Duration
}

func New() *gPScan {
	return &gPScan{}
}

func (s *gPScan) DoPortScanning() {
	allPortScanFlag, timeoutFlag := initiateFlags()
	args := parseFlags()
	argLength := len(args)

	if !isArgLengthValid(argLength) {
		fmt.Printf("Wrong count of arguments, expected 1-2, got %d\n", argLength)
		flag.Usage()
		return
	}

	var (
		protocol = args[0]
		port     int
	)

	s.protocol = protocol
	s.timeout = *timeoutFlag

	checkInputProtocolSupportedOrNot(protocol)

	if argLength == 2 {
		if !*allPortScanFlag {
			parsedPort, err := strconv.Atoi(args[1])
			port = parsedPort

			if err != nil {
				fmt.Println("wrong port format")
				return
			}
		} else {
			fmt.Println("'-a' flag is set, ignoring port")
		}

	}

	if !*allPortScanFlag {
		if s.IsPortOpen(port) {
			color.Green.Println("Port is open")
		} else {
			color.Danger.Println("Port is closed")
		}

		return
	}

	s.ScanAllPorts()
}

func (s *gPScan) ScanAllPorts() {
	var allOpenPorts []int
	for currentPort := 0; currentPort <= totalNumberOfPorts; currentPort++ {
		if s.IsPortOpen(currentPort) {
			allOpenPorts = append(allOpenPorts, currentPort)
		}
	}

	printAllOpenPorts(allOpenPorts)
}

func (s *gPScan) IsPortOpen(port int) bool {
	netAddress := fmt.Sprintf("localhost:%d", port)
	connection, err := net.DialTimeout(s.protocol, netAddress, s.timeout)
	if err != nil {
		return false
	}

	defer func() {
		err = connection.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("close connection (%s) : %w", netAddress, err))
		}
	}()

	return true
}

func printAllOpenPorts(allOpenPorts []int) {
	fmt.Print("Open ports: ")
	color.Green.Println(allOpenPorts)
}

func initiateFlags() (allPortsScanFlag *bool, timeoutFlag *time.Duration) {
	allPortsScanFlag = flag.Bool("a", false, "Scan all ports")
	timeoutFlag = flag.Duration("timeout", time.Millisecond*200, "Dialing timeout")
	return
}

func parseFlags() []string {
	flag.Parse()
	return flag.Args()
}

func isArgLengthValid(argLength int) bool {
	return argLength == 1 || argLength == 2
}

func checkInputProtocolSupportedOrNot(protocol string) {
	if _, ok := supportedProtocols[protocol]; !ok {
		fmt.Printf("protocol '%s' not supported\n", protocol)
		printAvailableProtocols()
		os.Exit(0)
	}
}

func printAvailableProtocols() {
	fmt.Print("available protocols: ")
	for protocol := range supportedProtocols {
		fmt.Print(protocol, ", ")
	}
	fmt.Println()
}
